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
	SetWidth(width int) Entry
	SetExportSelection(expsel bool) Entry
	SetJustify(lcr JustifyValue) Entry
	SetShow(s rune) Entry
	Value() string
	SetValue(v string)
}

type entry struct {
	*widget
	varname string
}

func (gt *GoTk) NewEntry(owner Widget) Entry {

	result := &entry{
		makeWidget(owner),
		randString(5),
	}

	result.instance.Send(fmt.Sprintf("ttk::entry %v -textvariable %v", result.path, result.varname))

	gt.widgetChannels[result.varname] = make(chan string)

	// todo - validate
	return result
}

func (e *entry) SetWidth(width int) Entry {
	widgetConfig(e, "width", width)
	return e
}

func (e *entry) SetExportSelection(expsel bool) Entry {
	widgetConfig(e, "exportselection", expsel)
	return e
}

func (e *entry) SetJustify(lcr JustifyValue) Entry {
	widgetConfig(e, "justify", lcr)
	return e
}

func (e *entry) SetShow(s rune) Entry {
	widgetConfig(e, "show", s)
	return e
}

func (e *entry) Value() string {
	return e.instance.sendAndGetResponse(e.varname, fmt.Sprintf("$::%v", e.varname), false)
}

func (e *entry) SetValue(v string) {
	e.instance.Send(fmt.Sprintf("set ::%v {%v}", e.varname, v))
}
