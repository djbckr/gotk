package gotk

import (
	"fmt"
	"strings"
)

/**************************************************************/

type configmap = map[string]interface{}

type GridConfig struct {
	inf      configmap
	itemPath string
	gt       *GoTk
}

func (gt *GoTk) StartGridConfig(slave Widget) *GridConfig {
	return &GridConfig{
		inf:      make(configmap),
		itemPath: slave.Path(),
		gt:       gt,
	}
}

func (g *GridConfig) Col(column int) *GridConfig {
	g.inf["column"] = column
	return g
}

func (g *GridConfig) ColSpan(span int) *GridConfig {
	g.inf["columnspan"] = span
	return g
}

func (g *GridConfig) In(master Widget) *GridConfig {
	g.inf["in"] = master.Path()
	return g
}

func (g *GridConfig) IPadX(amount int) *GridConfig {
	g.inf["ipadx"] = amount
	return g
}

func (g *GridConfig) IPadY(amount int) *GridConfig {
	g.inf["ipady"] = amount
	return g
}

func (g *GridConfig) PadX(amount int) *GridConfig {
	g.inf["padx"] = amount
	return g
}

func (g *GridConfig) PadY(amount int) *GridConfig {
	g.inf["pady"] = amount
	return g
}

func (g *GridConfig) Row(row int) *GridConfig {
	g.inf["row"] = row
	return g
}

func (g *GridConfig) RowSpan(span int) *GridConfig {
	g.inf["rowspan"] = span
	return g
}

func (g *GridConfig) Sticky(style string) *GridConfig {
	g.inf["sticky"] = style
	return g
}

func (g *GridConfig) Exec() {
	g.gt.Send(fmt.Sprintf("grid configure %v %v", g.itemPath, mapToString(g.inf)))
}

/**************************************************************/

type GridColRowConfig struct {
	inf      configmap
	itemPath string
	gt       *GoTk
	colrow   int
	which    string
}

// StartGridColConfig configures a grid column
// specify -1 for all columns
func (gt *GoTk) StartGridColConfig(master Widget, column int) *GridColRowConfig {
	return startGridColRowConfig(gt, master, column, "columnconfigure")
}

// StartGridRowConfig configures a grid row
// specify -1 for all columns
func (gt *GoTk) StartGridRowConfig(master Widget, row int) *GridColRowConfig {
	return startGridColRowConfig(gt, master, row, "rowconfigure")
}

func startGridColRowConfig(gt *GoTk, master Widget, colrow int, which string) *GridColRowConfig {
	return &GridColRowConfig{
		inf:      make(configmap),
		itemPath: master.Path(),
		gt:       gt,
		colrow:   colrow,
		which:    which,
	}
}

func (g *GridColRowConfig) MinSize(v int) *GridColRowConfig {
	g.inf["minsize"] = v
	return g
}

func (g *GridColRowConfig) Weight(v int) *GridColRowConfig {
	g.inf["weight"] = v
	return g
}

func (g *GridColRowConfig) Uniform(v string) *GridColRowConfig {
	g.inf["uniform"] = v
	return g
}

func (g *GridColRowConfig) Pad(v int) *GridColRowConfig {
	g.inf["pad"] = v
	return g
}

func (g *GridColRowConfig) Exec() {
	var colrow interface{} = g.colrow
	if g.colrow == -1 {
		colrow = "all"
	}
	g.gt.Send(fmt.Sprintf("grid %v %v %v %v", g.which, g.itemPath, colrow, mapToString(g.inf)))
}

/**************************************************************/
/**************************************************************/
/**************************************************************/
/**************************************************************/
/**************************************************************/
/**************************************************************/

func mapToString(theMap configmap) string {

	var sb strings.Builder

	var format string

	for k, v := range theMap {
		switch v.(type) {
		case string:
			format = " -%v {%v}"
		default:
			format = " -%v %v"
		}

		sb.WriteString(fmt.Sprintf(format, k, v))

	}

	return sb.String()
}
