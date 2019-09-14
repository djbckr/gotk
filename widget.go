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
	Children() []*widget
	Path() string
	getInstance() *GoTk
	getWidget() *widget
	// Parent() *widget
	addChild(child Widget)
	SetCursor(cursorType CursorType)
	SetTakeFocus(takeFocus bool)
}

type widget struct {
	instance *GoTk
	path     string
	parent   *widget
	children []*widget
}

func (w *widget) Children() []*widget {
	return w.children
}

// func (w *widget) Parent() *widget {
// 	return w.parent
// }

/*------------------------------------------------*/

func (w *widget) Path() string {
	return w.path
}

func (w *widget) SetCursor(cursorType CursorType) {
	widgetConfig(w, "cursor", cursorType)
}

func (w *widget) SetTakeFocus(takeFocus bool) {

	// don't do anything if this is a frame
	var f interface{} = w

	if _, found := f.(Frame); found == true {
		return
	}

	// everything else can take focus
	tf := 0

	if takeFocus {
		tf = 1
	}

	widgetConfig(w, "takefocus", tf)
}

/*------------------------------------------------*/

func (w *widget) getInstance() *GoTk {
	return w.instance
}

func (w *widget) getWidget() *widget {
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
	*widget
}

func (r *root) SetTitle(title string) {
	r.instance.Send(fmt.Sprintf("wm title . \"%v\"", title))
}
