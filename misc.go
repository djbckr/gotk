package gotk

import (
	"fmt"
	"strings"
)

type ModifierKey int

// Use these values or'd together when binding keystrokes on SetBindKey().
// If you aren't using one of these modifiers, pass 0 (zero)
const (
	CTRL  ModifierKey = 0x01
	ALT   ModifierKey = 0x02
	SHIFT ModifierKey = 0x04
	CMD   ModifierKey = 0x08
	DBL   ModifierKey = 0x10
	TRPL  ModifierKey = 0x20
)

// These constants are used for key-binding. To bind a normal key, like "a", use that.
const (
	Space     = "Space"
	BackSpace = "BackSpace"
	Tab       = "Tab"
	Return    = "Return"
	Escape    = "Escape"
	PgUp      = "Prior"
	PgDn      = "Next"
	End       = "End"
	Home      = "Home"
	Left      = "Left"
	Up        = "Up"
	Right     = "Right"
	Down      = "Down"
	Print     = "Print"
	Insert    = "Insert"
	Delete    = "Delete"
	Less      = "Less"
	Double    = "Double"
	Triple    = "Triple"
	F1        = "F1"
	F2        = "F2"
	F3        = "F3"
	F4        = "F4"
	F5        = "F5"
	F6        = "F6"
	F7        = "F7"
	F8        = "F8"
	F9        = "F9"
	F10       = "F10"
	F11       = "F11"
	F12       = "F12"
)

// Event is used to pass UI interactions back to your Go program. It is used in conjunction with EventChannel.
type Event interface {
	SourceWidget() Widget
	MouseCoordinates() (int, int)
	KeyPressed() string
}

type event struct {
	sourceWidget *widget
	x, y int
	key string
}

func (e *event) SourceWidget() Widget {
	return e.sourceWidget
}

func (e *event) MouseCoordinates() (x, y int) {
	return
}

func (e *event) KeyPressed() string {
	return e.key
}

/*
  EventChannel is what your Go program will listen to when you bind a key or button click.
  Your normal boilerplate might go something like this:

    mychannel := make(EventChannel) // optionally include buffer, if you feel like you might need it.

    go func() {
      for event := range mychannel {
        // do some processing here
      }
    }()

*/
type EventChannel = chan Event

func (gt *GoTk) SetBindKey(owner Widget, modifier ModifierKey, key string, eventChannel EventChannel) {

	var sb strings.Builder

	if len(key) > 1 {

		switch key {
		case Space, BackSpace, Tab, Return, Escape, PgUp, PgDn, End, Home, Left, Up, Right, Down, Print,
			Insert, Delete, F1, F2, F3, F4, F5, F6, F7, F8, F9, F10, F11, F12:
				sb.WriteString("-")
		    sb.WriteString(key)
		default:
			return
		}

	} else if len(key) == 1 {

		if key == " " {
			key = Space
		} else if key == "<" {
			key = Less
		}

		sb.WriteString("-")
		sb.WriteString(key)
	}

	vv := fmt.Sprintf("<%vKeyPress%v>", buildModifiers(modifier), sb.String())

	chName := randString(5)
	ch := make(chan string)

	widgetChannels[chName] = ch

	gt.Send(fmt.Sprintf("bind %v %v {puts $sockChan {¶%v¶%%A§%v§} ; flush $sockChan}", owner.Path(), vv, chName, chName))

	go func() {
		for {
			k := <-ch
			eventChannel <- &event{key: k}
		}
	}()

}

// SetFocus specifies the widget that should have focus.
func (gt *GoTk) SetFocus(widget Widget) {
	gt.Send(fmt.Sprintf("focus %v", widget.Path()))
}

func buildModifiers(modifier ModifierKey) string {
	var sb strings.Builder
	if modifier&CTRL > 0 {
		sb.WriteString("Control-")
	}
	if modifier&ALT > 0 {
		sb.WriteString("Alt-")
	}
	if modifier&SHIFT > 0 {
		sb.WriteString("Shift-")
	}
	if modifier&CMD > 0 {
		sb.WriteString("Command-")
	}
	if modifier&DBL > 0 {
		sb.WriteString("Double-")
	}
	if modifier&TRPL > 0 {
		sb.WriteString("Triple-")
	}
	return sb.String()
}
