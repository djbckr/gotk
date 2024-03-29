package gotk

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var alphabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var alphabetLen = len(alphabet)

type widgetChanType = map[string]chan string

// use our own rand/source so-as to not mess up somebody else's
var myRand = rand.New(rand.NewSource(time.Now().Unix()))

func randString(length int) string {
	var sb strings.Builder
	for i := 1; i <= length; i++ {
		idx := myRand.Int() % alphabetLen
		sb.WriteString(alphabet[idx])
	}
	return sb.String()
}

func makeName(parent Widget) string {
	n := randString(5)
	if parent.Path() != "." {
		n = "." + n
	}
	return parent.Path() + n
}

func makeWidget(owner Widget) *widget {
	result := &widget{
		instance: owner.getInstance(),
		path:     makeName(owner.getWidget()),
		parent:   owner.getWidget(),
	}
	owner.addChild(result)
	return result
}

func setPadding(w Widget, values ...int) {

	l := len(values) - 1

	var sb strings.Builder
	sb.WriteString("{")
	for i, v := range values {
		sb.WriteString(strconv.Itoa(v))
		if i < l {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("}")

	widgetConfig(w, "padding", sb.String())
}

func widgetConfig(w Widget, param string, value interface{}) {
	w.getInstance().Send(fmt.Sprintf("%v configure -%v %v", w.Path(), param, value))
}
