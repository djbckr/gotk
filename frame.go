package gotk

import "fmt"

type ReliefType = string

// Sets the type of border around a Frame. Used on SetRelief()
const (
	FLAT   ReliefType = "flat"
	GROOVE ReliefType = "groove"
	RAISED ReliefType = "raised"
	RIDGE  ReliefType = "ridge"
	SOLID  ReliefType = "solid"
	SUNKEN ReliefType = "sunken"
)

// Frame is a widget that contains other widgets. You typically use a frame to organize visual components.
type Frame interface {
	Widget
	SetPadding(...int) Frame
	SetWidth(width int) Frame
	SetHeight(height int) Frame
	SetBorderWidth(width int) Frame
	SetRelief(relief ReliefType) Frame
}

type frame struct {
	*widget
}

// NewFrame simply creates a new frame. You must provide the owner that this frame belongs to.
func (gt *GoTk) NewFrame(owner Widget) Frame {

	result := &frame{
		makeWidget(owner),
	}

	result.instance.Send(fmt.Sprintf("ttk::frame %v", result.path))

	return result
}

// SetPadding - will accept up to four integers
func (f *frame) SetPadding(values ...int) Frame {
	setPadding(f, values...)
	return f
}

// SetWidth sets the width of this frame in pixels. You should generally let the grid manager size the frame.
func (f *frame) SetWidth(width int) Frame {
	widgetConfig(f, "width", width)
	return f
}

// SetHeight sets the height of this frame in pixels. You should generally let the grid manager size the frame.
func (f *frame) SetHeight(height int) Frame {
	widgetConfig(f, "height", height)
	return f
}

func (f *frame) SetBorderWidth(width int) Frame {
	widgetConfig(f, "borderwidth", width)
	return f
}

func (f *frame) SetRelief(relief ReliefType) Frame {
	widgetConfig(f, "relief", relief)
	return f
}

