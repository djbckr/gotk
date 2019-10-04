package gotk

/*

   GoTk - a golang UI library using the Tcl/Tk library
   Copyright (C) 2019  Dean J. Becker

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU LESSER GENERAL PUBLIC LICENSE
   along with this program.  If not, see <https://www.gnu.org/licenses/>.

   The author can be reached at dean.becker@rubywillow.net

*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
)

// GoTk is the object you will interact with after you instantiate it with Tk()
type GoTk struct {
	cmd                *exec.Cmd
	stdIn              io.WriteCloser
	stdOut             io.ReadCloser
	stdErr             io.ReadCloser
	root               *root
	listener           net.Listener
	sendchan           chan string
	mouseWheelChName   string
	mouseWheelChannels []MouseWheelChannel
	widgetChannels     widgetChanType
	waitChannel        chan int
	dialogChName       string
}

// Tk instantiates a UI (TK) session. This is where you start.
func Tk() *GoTk {
	path, err := exec.LookPath("wish")
	if err != nil {
		log.Fatal("The \"wish\" command was not found on your system. Make sure tk is installed and in your path.")
	}

	cmd := exec.Command(path)
	pipeIn, _ := cmd.StdinPipe()
	pipeOut, _ := cmd.StdoutPipe()
	pipeErr, _ := cmd.StderrPipe()

	// stdout handler; anything here is preceded by a bullet (•) in the console
	go func() {

		scanner := bufio.NewScanner(pipeOut)
		for scanner.Scan() {
			fmt.Printf("• %v\n", scanner.Text())
		}

	}()

	// stderr handler - for now print warning. On OSX (my dev machine) something pops up that is just a non-event.
	go func() {

		scanner := bufio.NewScanner(pipeErr)
		for scanner.Scan() {
			fmt.Printf("WARNING FROM STDERR:\n%v\nEND WARNING FROM STDERR\n", scanner.Text())
		}

	}()

	// setup another stream for communication (listener) from wish to this application
	listener, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()

	if err != nil {
		log.Fatal("The \"wish\" command did not start properly!")
	}

	gotk := &GoTk{
		cmd:              cmd,
		stdIn:            pipeIn,
		stdOut:           pipeOut,
		stdErr:           pipeErr,
		listener:         listener,
		sendchan:         make(chan string, 10),
		mouseWheelChName: randString(5),
		widgetChannels:   make(widgetChanType),
		waitChannel:      make(chan int),
		dialogChName:     randString(5),
	}

	// setup the root window; it's already made for us by wish
	gotk.root = &root{
		&widget{
			instance: gotk,
			path:     ".",
		}}

	// all communication from wish to here is done in a goroutine
	go listenerControl(gotk)

	// all communication from here to wish is done in a goroutine
	go sender(gotk)

	// find the port that we opened on net.listener
	re := regexp.MustCompile(`\d+$`)
	port := re.FindString(listener.Addr().String())

	// tell wish to open a socket back to this program
	gotk.Send(fmt.Sprintf("set sockChan [socket localhost %v]", port))

	return gotk
}

// Wait for the main window to close
func (gt *GoTk) Wait() {
	<-gt.waitChannel
}

// Close tells TK to shutdown. This also shuts down our application.
func (gt *GoTk) Close() {
	gt.Send("exit")
}

// Root returns the root window.
func (gt *GoTk) Root() Root {
	return gt.root
}

// Send allows us to send any command to TK. For the most part, you will not need to use this as
// the components handle all the work for us. Sending a string here will append a new-line character.
// The command is sent in a go-routine and returns immediately if the channel is not full.
func (gt *GoTk) Send(s string) {
	gt.sendchan <- s
}

var reTest1 = regexp.MustCompile(`¶(\w+)¶`)
var reTest2 = regexp.MustCompile(`(?m)¶(\w+)¶((.|\n)*)§(\w+)§`)

// sender is the go-routine that sends commands to wish.
func sender(gt *GoTk) {
	for {
		msg := <-gt.sendchan
		ss := fmt.Sprintf("%v\n", msg)
		fmt.Printf("¶ %v", ss)
		_, _ = io.WriteString(gt.stdIn, ss)
	}
}

// listenerControl is the go-routine that receives data from wish
func listenerControl(gt *GoTk) {

	for {
		// this should only happen once....
		conn, err := gt.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// the data that comes through this pipe will be in the form of:
		// ¶varname¶some value that may have line feeds§varname§
		// if the value has line feeds, we won't have the §varname§ until the end
		// of the scan, so we have to build up the full response until we get §varname§
		// "varname" must be identical on both ends or that presents a problem.
		// once we have both ends, we can Send the contents via the channel of "varname"

		go func(c net.Conn) {
			scanner := bufio.NewScanner(c)
			var sb = strings.Builder{}

			for scanner.Scan() {

				if sb.Len() > 0 {
					sb.WriteString("\n")
				}

				sb.WriteString(scanner.Text())
				text := sb.String()

				fmt.Printf("» %v\n", text)

				rslt1 := reTest1.MatchString(text)
				rslt2 := reTest2.FindAllStringSubmatch(text, -1)

				if rslt1 && len(rslt2) > 0 {

					// we have a complete message

					if (len(rslt2[0]) != 5) || (rslt2[0][4] != rslt2[0][1]) {
						// a well-formed message will have rslt2[0] length of 5
						// and the capture of rslt2[0][4] should be the same as rslt2[0][1]
						// if not, we need to die - can't recover at this point.
						log.Fatal("Unexpected data on listening channel!")
					}

					// find the channel to pipe to
					ch := gt.widgetChannels[rslt2[0][1]]
					if ch != nil {
						ch <- rslt2[0][2]
					}
					sb.Reset()
				}

			}
			// If we get here, the pipe has been closed.
			// if program is Wait-ing, this unblocks the wait.
			gt.waitChannel <- 1
		}(conn)
	}
}

// sendAndGetResponse
// We need information from Tk. In order to do this, we have to tell it to send
// data on the $sockChan, but we also need to identify the variable/channel it belongs
// to.  We use ¶varname¶ for the header and §varname§ for the trailer, and whatever
// the command is in-between. Flushing $sockChan is required for timely response.
func (gt *GoTk) sendAndGetResponse(channelName string, command string, isSubCommand bool) (response string) {

	ch := gt.widgetChannels[channelName]

	// channel may not exist, which is fine; create it
	if ch == nil {
		ch = make(chan string)
		gt.widgetChannels[channelName] = ch
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("puts -nonewline $sockChan {¶%v¶} ; ", channelName))
	if isSubCommand {
		// a sub-command needs brackets around it
		sb.WriteString(fmt.Sprintf("puts -nonewline $sockChan [%v] ; ", command))
	} else {
		// this is likely a variable name, so no brackets here
		sb.WriteString(fmt.Sprintf("puts -nonewline $sockChan %v ; ", command))
	}
	sb.WriteString(fmt.Sprintf("puts $sockChan {§%v§} ; flush $sockChan", channelName))

	gt.Send(sb.String())

	// wait for message on the channel
	response = <-ch

	return
}
