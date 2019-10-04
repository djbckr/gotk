package gotk

import (
	"fmt"
	"strconv"
	"strings"
)

func (gt *GoTk) SetGridAnchor(widget Widget, anchor string) {
	gt.Send(fmt.Sprintf("grid anchor %v {%v}", widget.Path(), anchor))
}

func (gt *GoTk) SetGridBBox(widget Widget, top int, left int, width int, height int) {
	gt.Send(fmt.Sprintf("grid bbox %v %v %v %v %v", widget.Path(), top, left, width, height))
}

var gridBBoxChName = randString(5)

func (gt *GoTk) GridBBox(widget Widget) (top int, left int, width int, height int) {

	rslt := gt.sendAndGetResponse(gridBBoxChName, fmt.Sprintf("grid bbox %v", widget.Path()), true)

	r2 := strings.Split(rslt, " ")

	if len(r2) != 4 {
		return
	}

	top, _ = strconv.Atoi(r2[0])
	left, _ = strconv.Atoi(r2[1])
	width, _ = strconv.Atoi(r2[2])
	height, _ = strconv.Atoi(r2[3])

	return
}

type configmap = map[string]interface{}

// GridConfig is a structure that builds a message to send to the wish interpreter. You start with
// StartGridConfig() to get an instance of GridConfig, then call functions on it until you are done, then call Exec()
type GridConfig struct {
	inf      configmap
	itemPath string
	gt       *GoTk
}

// StartGridConfig - pass the widget you intend to place on the grid. This returns an instance of
// GridConfig after which you call all the functions you need, then call Exec() to finish the configuration.
func (gt *GoTk) StartGridConfig(widget Widget) *GridConfig {
	return &GridConfig{
		inf:      make(configmap),
		itemPath: widget.Path(),
		gt:       gt,
	}
}

// Col specifies the column of where to put the widget.
func (g *GridConfig) Col(column int) *GridConfig {
	g.inf["column"] = column
	return g
}

// ColSpan specifies that the widget spans multiple columns.
func (g *GridConfig) ColSpan(span int) *GridConfig {
	g.inf["columnspan"] = span
	return g
}

// In specifies the widget to insert this widget into. Generally you don't call this function.
func (g *GridConfig) In(master Widget) *GridConfig {
	g.inf["in"] = master.Path()
	return g
}

// IPadX - The amount specifies how much horizontal internal padding to leave on each side of the widget.
func (g *GridConfig) IPadX(amount int) *GridConfig {
	g.inf["ipadx"] = amount
	return g
}

// IPadY - The amount specifies how much vertical internal padding to leave on the top and bottom of the widget.
func (g *GridConfig) IPadY(amount int) *GridConfig {
	g.inf["ipady"] = amount
	return g
}

// PadX - The amount specifies how much horizontal external padding to leave on each side of the widget.
// Amount may be a list of two values to specify padding for left and right separately.
func (g *GridConfig) PadX(amount ...int) *GridConfig {
	g.inf["padx"] = intsToString(amount...)
	return g
}

// PadY - The amount specifies how much vertical external padding to leave on the top and bottom of the widget.
// Amount may be a list of two values to specify padding for top and bottom separately.
func (g *GridConfig) PadY(amount ...int) *GridConfig {
	g.inf["pady"] = intsToString(amount...)
	return g
}

// Row specifies the row of where to put the widget
func (g *GridConfig) Row(row int) *GridConfig {
	g.inf["row"] = row
	return g
}

// RowSpan specifies that the widget spans multiple rows.
func (g *GridConfig) RowSpan(span int) *GridConfig {
	g.inf["rowspan"] = span
	return g
}

// Sticky - If a widget's cell is larger than its requested dimensions, this option may be used to position
// (or stretch) the widget within its cell. Style is a string that contains zero or more of the
// characters "n", "s", "e" or "w". The string can optionally contains spaces or commas, but they are ignored.
// Each letter refers to a side (north, south, east, or west) that the widget will stick to. If both "n" and "s"
// (or "e" and "w") are specified, the widget will be stretched to fill the entire height (or width) of its cavity.
// The default is "", which causes the widget to be centered in its cavity, at its requested size.
func (g *GridConfig) Sticky(style string) *GridConfig {
	g.inf["sticky"] = style
	return g
}

// Exec - after calling all of the requested configuration functions, call Exec() to send the configuration to the UI
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

// StartGridColConfig configures a grid column.
// Specify -1 for all columns
func (gt *GoTk) StartGridColConfig(master Widget, column int) *GridColRowConfig {
	return startGridColRowConfig(gt, master, column, "columnconfigure")
}

// StartGridRowConfig configures a grid row.
// Specify -1 for all rows
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

func intsToString(i ...int) string {
	var sb strings.Builder

	for _, x := range i {
		sb.WriteString(fmt.Sprintf("%v ", x))
	}

	s := sb.String()
	return s[:len(s)-1]
}
