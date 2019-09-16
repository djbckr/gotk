package gotk

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var geometryChName = randString(5)
var geometryRe = regexp.MustCompile(`(?m)^(\d+)x(\d+)\+(-?\d+)\+(-?\d+)$`)

func (gt *GoTk) WmGeometry(wnd Widget) (width, height, left, top int)  {

	response := gt.sendAndGetResponse(geometryChName, fmt.Sprintf("wm geometry %v", wnd.Path()), true)

	values := geometryRe.FindAllStringSubmatch(response, -1)

	if len(values) != 1 && len(values[0]) != 5 {
		log.Fatal("something wrong with geometry query")
	}

	width, _ = strconv.Atoi(values[0][1])
	height, _ = strconv.Atoi(values[0][2])
	left, _ = strconv.Atoi(values[0][3])
	top, _ = strconv.Atoi(values[0][4])
	return

}

func (gt *GoTk) WmSetGeometry(wnd Widget, width, height, left, top int) {
	gt.Send(fmt.Sprintf("wm geometry %v {%vx%v+%v+%v}", wnd.Path(), width, height, left, top))
}
