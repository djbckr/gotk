package gotk

import (
	"fmt"
	"log"
	"strconv"
)

type MouseWheelChannel = chan int

func (gt *GoTk) BindMouseWheel(owner Widget, key ModifierKey, mwChannel MouseWheelChannel)  {

	var ch chan string

	if ch = widgetChannels[gt.mouseWheelChName]; ch == nil {
		ch = make(chan string)
		widgetChannels[gt.mouseWheelChName] = ch
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
				log.Fatal("Problem reading mouse wheel")
			}

			for _, c := range gt.mouseWheelChannels {
				c <- v
			}
		}
	}()

}
