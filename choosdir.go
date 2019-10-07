package gotk

import (
	"fmt"
	"strings"
)

type ChooseDirectoryDlg interface {
	SetInitialDir(initialDir string) ChooseDirectoryDlg
	SetMustExist(mustExist bool) ChooseDirectoryDlg
	SetParent(parent Widget) ChooseDirectoryDlg
	SetTitle(title string) ChooseDirectoryDlg
	Exec() string
}

type choosedirectorydlg struct {
	initialDir string
	mustExist  bool
	parent     string
	title      string
	instance   *GoTk
}

func (gt *GoTk) ChooseDirectoryStart() ChooseDirectoryDlg {
	return &choosedirectorydlg{
		instance: gt,
	}
}

func (chdr *choosedirectorydlg) SetInitialDir(initialDir string) ChooseDirectoryDlg {
	chdr.initialDir = initialDir
	return chdr
}

func (chdr *choosedirectorydlg) SetMustExist(mustExist bool) ChooseDirectoryDlg {
	chdr.mustExist = mustExist
	return chdr
}

func (chdr *choosedirectorydlg) SetParent(parent Widget) ChooseDirectoryDlg {
	chdr.parent = parent.Path()
	return chdr
}

func (chdr *choosedirectorydlg) SetTitle(title string) ChooseDirectoryDlg {
	chdr.title = title
	return chdr
}

func (chdr *choosedirectorydlg) Exec() string {

	var sb strings.Builder

	sb.WriteString("set dlgResult [tk_chooseDirectory")

	if chdr.initialDir != "" {
		sb.WriteString(fmt.Sprintf(" -initialdir {%v}", chdr.initialDir))
	}

	if chdr.mustExist {
		sb.WriteString(fmt.Sprintf(" -mustexist %v", chdr.mustExist))
	}

	if chdr.parent != "" {
		sb.WriteString(fmt.Sprintf(" -parent {%v}", chdr.parent))
	}

	if chdr.title != "" {
		sb.WriteString(fmt.Sprintf(" -title {%v}", chdr.title))
	}

	sb.WriteString("]")
	chdr.instance.Send(sb.String())

	return chdr.instance.sendAndGetResponse(chdr.instance.dialogChName, "$::dlgResult", false)

}
