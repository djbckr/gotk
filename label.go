package gotk

import (
	"fmt"
	"log"
)

type Anchor = string

// These values are used to set the label anchor. These are the only allowable values.
const (
	N  Anchor = "n"
	NE Anchor = "ne"
	E  Anchor = "e"
	SE Anchor = "se"
	S  Anchor = "s"
	SW Anchor = "sw"
	W  Anchor = "w"
	NW Anchor = "nw"
)

type Label interface {
	Widget
	SetText(text string)
	SetWidth(width int) *label
	SetUnderline(underline int) *label
	SetRelief(reliefType ReliefType) *label
	SetPadding(values ...int) *label
	SetJustify(value JustifyValue) *label
	SetAnchor(anchor Anchor) *label
	SetBackground() *label
	SetFont(fontName string) *label
	SetForeground() *label
	SetWrapLength(length int) *label
}

type label struct {
	*widget
	varname string
}

func (gt *GoTk) NewLabel(owner Widget, text string) *label {

	result := &label{
		makeWidget(owner),
		randString(5),
	}

	owner.addChild(result)

	result.instance.Send(fmt.Sprintf("ttk::label %v -textvariable %v", result.path, result.varname))

	result.SetText(text)

	return result

}

func (l *label) SetText(text string) {
	l.instance.Send(fmt.Sprintf("set ::%v {%v}", l.varname, text))
}

func (l *label) SetWidth(width int) *label {
	widgetConfig(l, "width", width)
	return l
}

func (l *label) SetUnderline(underline int) *label {
	widgetConfig(l, "underline", underline)
	return l
}

func (l *label) SetRelief(reliefType ReliefType) *label {
	widgetConfig(l, "relief", reliefType)
	return l
}

func (l *label) SetPadding(values ...int) *label {
	setPadding(l, values...)
	return l
}

func (l *label) SetJustify(value JustifyValue) *label {
	widgetConfig(l, "justify", value)
	return l
}

func (l *label) SetAnchor(anchor Anchor) *label {
	switch anchor {
	case N, NE, E, SE, S, SW, W, NW, CENTER:
		// all good, do nothing
	default:
		log.Fatal("Invalid value passed to Label.SetAnchor()")
	}
	widgetConfig(l, "anchor", anchor)
	return l
}

func (l *label) SetBackground() *label {
	// todo - implement
	return l
}

func (l *label) SetFont(fontName string) *label {
	widgetConfig(l, "font", "{"+fontName+"}")
	return l
}

func (l *label) SetForeground() *label {
	// todo -implement
	return l
}

func (l *label) SetWrapLength(length int) *label {
	widgetConfig(l, "wraplength", length)
	return l
}
