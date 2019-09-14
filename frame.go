package gotk

import "fmt"

type ReliefType = string

const (
	FLAT   ReliefType = "flat"
	GROOVE ReliefType = "groove"
	RAISED ReliefType = "raised"
	RIDGE  ReliefType = "ridge"
	SOLID  ReliefType = "solid"
	SUNKEN ReliefType = "sunken"
)

type Frame interface {
	Widget
	SetPadding(...int) *frame
	SetWidth(width int) *frame
	SetHeight(height int) *frame
	SetBorderWidth(width int) *frame
	SetRelief(relief ReliefType) *frame
}

type frame struct {
	*widget
}

func (gt *GoTk) NewFrame(owner Widget) *frame {

	result := &frame{
		makeWidget(owner),
	}

	owner.addChild(result)

	result.instance.Send(fmt.Sprintf("ttk::frame %v", result.path))

	return result
}

func (f *frame) SetPadding(values ...int) *frame {
	setPadding(f, values...)
	return f
}

func (f *frame) SetWidth(width int) *frame {
	widgetConfig(f, "width", width)
	return f
}

func (f *frame) SetHeight(height int) *frame {
	widgetConfig(f, "height", height)
	return f
}

func (f *frame) SetBorderWidth(width int) *frame {
	widgetConfig(f, "borderwidth", width)
	return f
}

func (f *frame) SetRelief(relief ReliefType) *frame {
	widgetConfig(f, "relief", relief)
	return f
}

