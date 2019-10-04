package gotk

import (
	"fmt"
	"log"
	"strconv"
)

// MouseWheelChannel is a simple channel that passes integers when the mouse-wheel is spun. See BindMouseWheel().
type MouseWheelChannel = chan int

// BindMouseWheel will send mwChannel a bunch of integers when the mouse wheel is spun.
// If you call this more than once, it will send the same message to multiple channels.
func (gt *GoTk) BindMouseWheel(owner Widget, key ModifierKey, mwChannel MouseWheelChannel)  {

	var ch chan string

	if ch = gt.widgetChannels[gt.mouseWheelChName]; ch == nil {
		ch = make(chan string)
		gt.widgetChannels[gt.mouseWheelChName] = ch
	}

	gt.mouseWheelChannels = append(gt.mouseWheelChannels, mwChannel)

	if len(gt.mouseWheelChannels) > 1 {
		return
	}

	gt.Send(fmt.Sprintf("bind %v <%vMouseWheel> {puts $sockChan {¶%v¶%%D§%v§} ; flush $sockChan}",
		owner.Path(), buildModifiers(key), gt.mouseWheelChName, gt.mouseWheelChName))

	go func() {
		for {
			s := <- ch
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Problem reading mouse wheel", err)
			}

			for _, c := range gt.mouseWheelChannels {
				c <- v
			}
		}
	}()

}
