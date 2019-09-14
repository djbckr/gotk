package gotk

import "fmt"

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
	fnName string
}

func (gt *GoTk) NewButton(owner Widget, text string, fn func(btn Button)) *button {

	result := &button{
		makeWidget(owner),
		randString(5),
	}

	owner.addChild(result)

	widgetChannels[result.fnName] = make(chan string)

	result.instance.Send(fmt.Sprintf("proc %v {} {", result.fnName))
	result.instance.Send(fmt.Sprintf("  global sockChan"))
	result.instance.Send(fmt.Sprintf("  puts $sockChan {¶%v¶x§%v§}", result.fnName, result.fnName))
	result.instance.Send(fmt.Sprintf("  flush $sockChan"))
	result.instance.Send(fmt.Sprintf("}"))

	result.instance.Send(fmt.Sprintf("ttk::button %v -text {%v} -command %v", result.path, text, result.fnName))

	go func() {
		for {
			<- widgetChannels[result.fnName]
			fn(result)
		}
	}()

	return result
}

func (b *button) SetState(state WidgetState) *button {
	widgetConfig(b, "state", state)
	return b
}

func (b *button) SetText(text string) *button {
	widgetConfig(b, "text", "{" + text + "}")
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
	return b.fnName
}


