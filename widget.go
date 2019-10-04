package gotk

import (
	"fmt"
)

/* let's not forget value.(type) which will return the interface the value implements */
/*********************************************************************/

type CursorType = string

const (
	ARROW     CursorType = "arrow"
	CROSSHAIR CursorType = "crosshair"
	IBEAM     CursorType = "ibeam"
	NONE      CursorType = "none"
	WATCH     CursorType = "watch"
	XTERM     CursorType = "xterm"
)

type Widget interface {
	Children() []Widget
	Path() string
	getInstance() *GoTk
	getWidget() Widget
	// Parent() Widget
	addChild(child Widget)
	SetCursor(cursorType CursorType) Widget
	SetTakeFocus(takeFocus bool) Widget
	SetState(state WidgetState) Widget
}

type widget struct {
	instance *GoTk
	path     string
	parent   Widget
	children []Widget
}

func (w *widget) Children() []Widget {
	return w.children
}

// func (w *widget) Parent() Widget {
// 	return w.parent
// }

/*------------------------------------------------*/

func (w *widget) Path() string {
	return w.path
}

func (w *widget) SetState(state WidgetState) Widget {
	widgetConfig(w, "state", state)
	return w
}

func (w *widget) SetCursor(cursorType CursorType) Widget {
	widgetConfig(w, "cursor", cursorType)
	return w
}

func (w *widget) SetTakeFocus(takeFocus bool) Widget {

	// don't do anything if this is a frame
	var f interface{} = w

	if _, found := f.(Frame); found == true {
		return w
	}

	// everything else can take focus
	tf := 0

	if takeFocus {
		tf = 1
	}

	widgetConfig(w, "takefocus", tf)
	return w
}

/*------------------------------------------------*/

func (w *widget) getInstance() *GoTk {
	return w.instance
}

func (w *widget) getWidget() Widget {
	return w
}

func (w *widget) addChild(child Widget) {
	w.children = append(w.children, child.getWidget())
}

/*********************************************************************/

type Root interface {
	Widget
	SetTitle(title string)
}

type root struct {
	Widget
}

func (r *root) SetTitle(title string) {
	r.getInstance().Send(fmt.Sprintf("wm title . \"%v\"", title))
}
