package gotk

import (
	"fmt"
)

type WidgetState = string

// Values used to set the state of a widget.
const (
	NORMAL   WidgetState = "normal"
	DISABLED WidgetState = "disabled"
	READONLY WidgetState = "readonly"
)

type JustifyValue = string

// Values used to set the alignment of a widget.
const (
	LEFT   JustifyValue = "left"
	CENTER JustifyValue = "center"
	RIGHT  JustifyValue = "right"
)

// This is a text input box.
type Entry interface {
	Widget
	SetWidth(width int) *entry
	SetExportSelection(expsel bool) *entry
	SetJustify(lcr JustifyValue) *entry
	SetShow(s rune) *entry
	SetState(state WidgetState) *entry
	Value() string
	SetValue(v string)
}

type entry struct {
	*widget
	varname string
}

func (gt *GoTk) NewEntry(owner Widget) *entry {

	result := &entry{
		makeWidget(owner),
		randString(5),
	}

	owner.addChild(result)

	result.instance.Send(fmt.Sprintf("ttk::entry %v -textvariable %v", result.path, result.varname))

	widgetChannels[result.varname] = make(chan string)

	// todo - validate
	return result
}

func (e *entry) SetWidth(width int) *entry {
	widgetConfig(e, "width", width)
	return e
}

func (e *entry) SetExportSelection(expsel bool) *entry {
	widgetConfig(e, "exportselection", expsel)
	return e
}

func (e *entry) SetJustify(lcr JustifyValue) *entry {
	widgetConfig(e, "justify", lcr)
	return e
}

func (e *entry) SetShow(s rune) *entry {
	widgetConfig(e, "show", s)
	return e
}

func (e *entry) SetState(state WidgetState) *entry {
	widgetConfig(e, "state", state)
	return e
}

func (e *entry) Value() string {
	return e.instance.sendAndGetResponse(e.varname, fmt.Sprintf("$::%v", e.varname), false)
}

func (e *entry) SetValue(v string) {
	e.instance.Send(fmt.Sprintf("set ::%v {%v}", e.varname, v))
}
