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
	SetWidth(width int) Label
	SetUnderline(underline int) Label
	SetRelief(reliefType ReliefType) Label
	SetPadding(values ...int) Label
	SetJustify(value JustifyValue) Label
	SetAnchor(anchor Anchor) Label
	SetBackground() Label
	SetFont(fontName string) Label
	SetForeground() Label
	SetWrapLength(length int) Label
}

type label struct {
	*widget
	varname string
}

func (gt *GoTk) NewLabel(owner Widget, text string) Label {

	result := &label{
		makeWidget(owner),
		randString(5),
	}

	result.instance.Send(fmt.Sprintf("ttk::label %v -textvariable %v", result.path, result.varname))

	result.SetText(text)

	return result

}

func (l *label) SetText(text string) {
	l.instance.Send(fmt.Sprintf("set ::%v {%v}", l.varname, text))
}

func (l *label) SetWidth(width int) Label {
	widgetConfig(l, "width", width)
	return l
}

func (l *label) SetUnderline(underline int) Label {
	widgetConfig(l, "underline", underline)
	return l
}

func (l *label) SetRelief(reliefType ReliefType) Label {
	widgetConfig(l, "relief", reliefType)
	return l
}

func (l *label) SetPadding(values ...int) Label {
	setPadding(l, values...)
	return l
}

func (l *label) SetJustify(value JustifyValue) Label {
	widgetConfig(l, "justify", value)
	return l
}

func (l *label) SetAnchor(anchor Anchor) Label {
	switch anchor {
	case N, NE, E, SE, S, SW, W, NW, CENTER:
		// all good, do nothing
	default:
		log.Fatal("Invalid value passed to Label.SetAnchor()")
	}
	widgetConfig(l, "anchor", anchor)
	return l
}

func (l *label) SetBackground() Label {
	// todo - implement
	return l
}

func (l *label) SetFont(fontName string) Label {
	widgetConfig(l, "font", "{"+fontName+"}")
	return l
}

func (l *label) SetForeground() Label {
	// todo -implement
	return l
}

func (l *label) SetWrapLength(length int) Label {
	widgetConfig(l, "wraplength", length)
	return l
}
