package gotk

import (
	"fmt"
)

type Button interface {
	Widget
	SetState(state WidgetState) *button
	SetText(text string) *button
	SetUnderline(underline int) *button
	SetWidth(width int) *button
	getFnName() string
	// SetDefault(dflt bool) *button
}

type button struct {
	*widget
	callbackId string
}

func (gt *GoTk) NewButton(owner Widget, text string, channel EventChannel) *button {

	result := &button{
		makeWidget(owner),
		randString(5),
	}

	owner.addChild(result)

	widgetChannels[result.callbackId] = make(chan string)

	// result.instance.Send(fmt.Sprintf("proc %v {} {", result.callbackId))
	// result.instance.Send(fmt.Sprintf("  global sockChan"))
	// result.instance.Send(fmt.Sprintf("  puts $sockChan {¶%v¶x§%v§}", result.callbackId, result.callbackId))
	// result.instance.Send(fmt.Sprintf("  flush $sockChan"))
	// result.instance.Send(fmt.Sprintf("}"))

	result.instance.Send(fmt.Sprintf("ttk::button %v -text {%v} -command {puts $sockChan {¶%v¶§%v§} ; flush $sockChan}", result.path, text, result.callbackId, result.callbackId))

	go func() {
		for {
			<-widgetChannels[result.callbackId]
			channel <- &event{
				sourceWidget: result.widget,
			}
		}
	}()

	return result
}

func (b *button) SetState(state WidgetState) *button {
	widgetConfig(b, "state", state)
	return b
}

func (b *button) SetText(text string) *button {
	widgetConfig(b, "text", "{"+text+"}")
	return b
}

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
