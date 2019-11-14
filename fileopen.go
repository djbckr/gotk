package gotk

import (
	"fmt"
	"strings"
)

// This is a structure that you use to determine what file types may be shown in the FileOpenDlg.
// For example:
//    fileOpenType := FileOpenType{"PDF Files", ".pdf"}
type FileOpenType struct {
	Description string
	Pattern     string
}

type FileOpenTypes = []FileOpenType

// FileOpenDlg lets the user to choose one or more files. You can optionally call the configuration functions
// and finally call Exec() to display the dialog and allow the user to make their selection.
type FileOpenDlg interface {

	// (Optional) Attach this dialog to a widget (window)
	SetParent(widget Widget) FileOpenDlg

	// (Optional) Pass a slice of FileOpenTypes to limit the selection that the user can choose.
	SetFileTypes(types FileOpenTypes) FileOpenDlg

	// (Optional) Open the dialog in this directory
	SetInitialDir(dir string) FileOpenDlg

	// (Optional) Set the selected file to this file.
	SetInitialFile(file string) FileOpenDlg

	// (Optional) Allow the user to select multiple files. Default is one file.
	SetMultiple(multiple bool) FileOpenDlg

	// (Optional) Set the title of this dialog
	SetTitle(title string) FileOpenDlg

	// Display the dialog to the user. Regardless of whether you have set multiple selections,
	// this always returns an array of strings. If the user made a selection, the array will
	// have one or more strings. If the user canceled the dialog, you will receive an empty array.
	Exec() (results []string)
}

type fileopendlg struct {
	parent      string
	fileTypes   FileOpenTypes
	initialDir  string
	initialFile string
	multiple    bool
	title       string
	instance    *GoTk
}

func (gt *GoTk) FileOpenStart() FileOpenDlg {
	return &fileopendlg{
		instance: gt,
	}
}

func (fo *fileopendlg) SetParent(widget Widget) FileOpenDlg {
	fo.parent = widget.Path()
	return fo
}

func (fo *fileopendlg) SetFileTypes(types FileOpenTypes) FileOpenDlg {
	fo.fileTypes = types
	return fo
}

func (fo *fileopendlg) SetInitialDir(dir string) FileOpenDlg {
	fo.initialDir = dir
	return fo
}

func (fo *fileopendlg) SetInitialFile(file string) FileOpenDlg {
	fo.initialFile = file
	return fo
}

func (fo *fileopendlg) SetMultiple(multiple bool) FileOpenDlg {
	fo.multiple = multiple
	return fo
}

func (fo *fileopendlg) SetTitle(title string) FileOpenDlg {
	fo.title = title
	return fo
}

func (fo *fileopendlg) Exec() (results []string) {
	var sb strings.Builder

	sb.WriteString("set dlgResult [tk_getOpenFile")

	if fo.multiple {
		sb.WriteString(" -multiple true")
	}

	if fo.parent != "" {
		sb.WriteString(" -parent ")
		sb.WriteString(fo.parent)
	}

	if fo.fileTypes != nil {
		sb.WriteString(" -filetypes ")
		sb.WriteString(formatFileTypes(fo.fileTypes))
	}

	if fo.initialFile != "" {
		sb.WriteString(fmt.Sprintf(" -initialfile {%v}", fo.initialFile))
	}

	if fo.initialDir != "" {
		sb.WriteString(fmt.Sprintf(" -initialdir {%v}", fo.initialDir))
	}

	if fo.title != "" {
		sb.WriteString(fmt.Sprintf(" -title {%v}", fo.title))
	}

	sb.WriteString("]")
	fo.instance.Send(sb.String())

	result := fo.instance.sendAndGetResponse(fo.instance.dialogChName, "$::dlgResult", false)

	return parseFileNameList(result)

}

func formatFileTypes(ft FileOpenTypes) string {
	var sb strings.Builder

	sb.WriteString("{ ")

	for _, v := range ft {
		sb.WriteString(fmt.Sprintf("{ {%v} {%v} } ", v.Description, v.Pattern))
	}

	sb.WriteString("}")

	return sb.String()
}

// TK has a pretty bizarre way of returning the selection from this dialog.
// First, it's a single string that is returned even with multiple selections.
// Each result is separated by one space. If a filename has a space, TK wraps
// the filename in {}. The filename could have brace chars {} within, so we
// have to watch for that as well. When it needs to, it'll prefix a literal
// with a backslash. (I don't know what this does on Windows)
func parseFileNameList(src string) (result []string) {

	if src == "" {
		return
	}

	var chr rune
	// a level > 0 means the string is wrapped in {} and we need to look for }
	// a level < 0 means it's a normal string and we need to look for a space.
	var level = 0
	var nextLiteral bool

	var sb strings.Builder

	for _, chr = range src {

		if nextLiteral == false && chr == '\\' {
			nextLiteral = true
			continue
		}

		if nextLiteral {
			sb.WriteRune(chr)
			nextLiteral = false
			continue
		}

		if level >= 0 && chr == '{' {
			level++
			if level == 1 {
				continue
			}
		} else if level == 0 && chr != ' ' {
			level = -1
		} else if level == 0 {
			continue
		}

		if level >= 0 && chr == '}' {
			level--
		}

		if level < 0 && chr == ' ' {
			level = 0
		}

		if level != 0 {
			sb.WriteRune(chr)
		} else {
			result = append(result, sb.String())
			sb.Reset()
		}

	}

	result = append(result, sb.String())

	return
}
