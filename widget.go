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

type WidgetState = string

// Values used to set the state of a widget.
const (
	NORMAL   WidgetState = "normal"
	DISABLED WidgetState = "disabled"
	READONLY WidgetState = "readonly"
)

type Widget interface {

	// Children returns the widgets this widget owns
	Children() []Widget

	// Path returns the widget name used in the wish interpreter
	Path() string
	getInstance() *GoTk
	getWidget() Widget
	// Parent() Widget
	addChild(child Widget)

	// SetCursor sets the CursorType for this widget
	SetCursor(cursorType CursorType) Widget

	// SetTakeFocus tells TK to allow this widget to focus when tabbing through components
	SetTakeFocus(takeFocus bool) Widget

	// SetState sets the widget to disabled, readonly, or normal
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
