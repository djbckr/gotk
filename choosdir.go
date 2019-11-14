package gotk

import (
	"fmt"
	"strings"
)

// The ChooseDirectoryDlg interface defines the optional functions you can call to configure the dialog.
// Call Exec() to display the dialog and return the chosen directory.
type ChooseDirectoryDlg interface {

	// (Optional) Start the directory search here.
	SetInitialDir(initialDir string) ChooseDirectoryDlg

	// (Optional) The directory must exist before selection is complete.
	SetMustExist(mustExist bool) ChooseDirectoryDlg

	// (Optional) The widget (window) that this dialog is attached to.
	SetParent(parent Widget) ChooseDirectoryDlg

	// (Optional) The title of the dialog window.
	SetTitle(title string) ChooseDirectoryDlg

	// When finished optionally calling the above functions, call this function to display the dialog. It will return
	// a string of the chosen directory. If the string is empty, the user canceled the dialog.
	Exec() string
}

type choosedirectorydlg struct {
	initialDir string
	mustExist  bool
	parent     string
	title      string
	instance   *GoTk
}

// Startup a ChooseDirectoryDlg instance. Call the routines on the returned interface to configure
// the dialog, then call .Exec() to display the dialog.
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
