package gotk

import (
	"fmt"
)

type JustifyValue = string

// Values used to set the alignment of a widget.
const (
	LEFT   JustifyValue = "left"
	CENTER JustifyValue = "center"
	RIGHT  JustifyValue = "right"
)

// This is a text input box. Use this to allow the user to enter some text to the UI.
type Entry interface {
	Widget

	// (Optional) Set the width of the input box.
	SetWidth(width int) Entry

	// (Optional) Set the alignment of the text within the input box
	SetJustify(lcr JustifyValue) Entry

	// (Optional) If this is intended to be used as a password box,
	// set the character to be shown as the placeholder for the underlying text.
	SetShow(s rune) Entry

	// Get the current value of this Entry
	Value() string

	// Set the value of this Entry
	SetValue(v string)
}

type entry struct {
	*widget
	varname string
}

// Create an entry box to eventually be placed on the screen. The owner is the parent widget that
// this Entry will belong to.
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
