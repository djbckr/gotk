package gotk

import (
	"fmt"
	"strings"
)

type FileSaveDlg interface {
	SetConfirmOverwrite(confirmOverwrite bool) FileSaveDlg
	SetDefaultExtension(defaultExtension string) FileSaveDlg
	SetInitialDir(initialDir string) FileSaveDlg
	SetInitialFile(initialFile string) FileSaveDlg
	SetParent(parent Widget) FileSaveDlg
	SetTitle(title string) FileSaveDlg
	Exec() string
}

type filesavedlg struct {
	confirmOverwrite bool
	defaultExtension string
	initialDir       string
	initialFile      string
	parent           string
	title            string
	instance         *GoTk
}

func (gt *GoTk) FileSaveStart() FileSaveDlg {
	return &filesavedlg{
		confirmOverwrite: true,
		instance:         gt,
	}
}

func (fs *filesavedlg) SetConfirmOverwrite(confirmOverwrite bool) FileSaveDlg {
	fs.confirmOverwrite = confirmOverwrite
	return fs
}

func (fs *filesavedlg) SetDefaultExtension(defaultExtension string) FileSaveDlg {
	fs.defaultExtension = defaultExtension
	return fs
}

func (fs *filesavedlg) SetInitialDir(initialDir string) FileSaveDlg {
	fs.initialDir = initialDir
	return fs
}

func (fs *filesavedlg) SetInitialFile(initialFile string) FileSaveDlg {
	fs.initialFile = initialFile
	return fs
}

func (fs *filesavedlg) SetParent(parent Widget) FileSaveDlg {
	fs.parent = parent.Path()
	return fs
}

func (fs *filesavedlg) SetTitle(title string) FileSaveDlg {
	fs.title = title
	return fs
}

func (fs *filesavedlg) Exec() string {
	var sb strings.Builder

	sb.WriteString("set dlgResult [tk_getSaveFile")

	sb.WriteString(fmt.Sprintf(" -confirmoverwrite %v", fs.confirmOverwrite))

	if fs.defaultExtension != "" {
		sb.WriteString(fmt.Sprintf(" -defaultextension {%v}", fs.defaultExtension))
	}

	if fs.initialDir != "" {
		sb.WriteString(fmt.Sprintf(" -initialdir {%v}", fs.initialDir))
	}

	if fs.initialFile != "" {
		sb.WriteString(fmt.Sprintf(" -initialfile {%v}", fs.initialFile))
	}

	if fs.parent != "" {
		sb.WriteString(fmt.Sprintf(" -parent {%v}", fs.parent))
	}

	if fs.title != "" {
		sb.WriteString(fmt.Sprintf(" -title {%v}", fs.title))
	}

	sb.WriteString("]")
	fs.instance.Send(sb.String())

	return fs.instance.sendAndGetResponse(fs.instance.dialogChName, "$::dlgResult", false)

}
