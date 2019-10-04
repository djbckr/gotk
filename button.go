package gotk

import (
	"fmt"
)

type ButtonState = string

const (
	ButtonStateNormal ButtonState = "normal"
	ButtonStateActive ButtonState = "active"
	ButtonStateDisabled ButtonState = "disabled"
)

// Button is a widget that is "clickable".
type Button interface {
	Widget
	SetText(text string) *button
	SetUnderline(underline int) *button
	SetWidth(width int) *button
	SetDefault(state ButtonState) *button
	getFnName() string
}

type button struct {
	*widget
	callbackId string
}

// Create a new button inside owner. You must provide and listen to an EventChannel to receive clicks.
func (gt *GoTk) NewButton(owner Widget, text string, channel EventChannel) *button {

	result := &button{
		makeWidget(owner),
		randString(5),
	}

	gt.widgetChannels[result.callbackId] = make(chan string)

	result.instance.Send(fmt.Sprintf("ttk::button %v -text {%v} -command {puts $sockChan {¶%v¶§%v§} ; flush $sockChan}", result.path, text, result.callbackId, result.callbackId))

	go func() {
		for {
			<-gt.widgetChannels[result.callbackId]
			channel <- &event{
				sourceWidget: result.widget,
			}
		}
	}()

	return result
}

// SetDefault allows the following values: ButtonStateNormal, ButtonStateActive, ButtonStateDisabled
func (b *button) SetDefault(state ButtonState) *button {
	widgetConfig(b, "default", state)
	return b
}

// SetText within the button
func (b *button) SetText(text string) *button {
	widgetConfig(b, "text", "{"+text+"}")
	return b
}

// SetUnderline specifies the integer index (0-based) of a character to underline in the text string. The underlined character is used for mnemonic activation.
func (b *button) SetUnderline(underline int) *button {
	widgetConfig(b, "underline", underline)
	return b
}

func (b *button) SetWidth(width int) *button {
	widgetConfig(b, "width", width)
	return b
}

func (b *button) getFnName() string {
	return b.callbackId
}
