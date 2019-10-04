package gotk

import (
	"fmt"
	"strings"
)

type FileOpenType struct {
	Description string
	Pattern     string
}

type FileOpenTypes = []FileOpenType

type FileOpen interface {
	SetParent(widget Widget) *fileopen
	SetFileTypes(types FileOpenTypes) *fileopen
	SetInitialDir(dir string) *fileopen
	SetInitialFile(file string) *fileopen
	SetMultiple(multiple bool) *fileopen
	SetTitle(title string) *fileopen
	Exec() (results []string, err error)
}

type fileopen struct {
	parent      string
	fileTypes   FileOpenTypes
	initialDir  string
	initialFile string
	multiple    bool
	title       string
	instance    *GoTk
}

func (gt *GoTk) FileOpenStart() *fileopen {
	return &fileopen{
		instance: gt,
	}
}

func (fo *fileopen) SetParent(widget Widget) *fileopen {
	fo.parent = widget.Path()
	return fo
}

func (fo *fileopen) SetFileTypes(types FileOpenTypes) *fileopen {
	fo.fileTypes = types
	return fo
}

func (fo *fileopen) SetInitialDir(dir string) *fileopen {
	fo.initialDir = dir
	return fo
}

func (fo *fileopen) SetInitialFile(file string) *fileopen {
	fo.initialFile = file
	return fo
}

func (fo *fileopen) SetMultiple(multiple bool) *fileopen {
	fo.multiple = multiple
	return fo
}

func (fo *fileopen) SetTitle(title string) *fileopen {
	fo.title = title
	return fo
}

func (fo *fileopen) Exec() (results []string, err error) {
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

func parseFileNameList(src string) (result []string, err error) {

	if src == "" {
		return
	}

	var chr rune
	var level = 0

	var sb strings.Builder

	for _, chr = range src {

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
