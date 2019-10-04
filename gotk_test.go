package gotk_test

import (
	"fmt"
	"github.com/djbckr/gotk"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestTk(t *testing.T) {
	// instantiate Tk (executes the wish program)
	ui := gotk.Tk()

	// Tk() also instantiates the root window, so let's get that
	root := ui.Root()

	// The following covers the script found at https://tkdocs.com/tutorial/firstexample.html#code

	// wm title . "Feet to Meters"
	root.SetTitle("Feet to Meters")

	// grid [ttk::frame .c -padding "3 3 12 12"] -column 0 -row 0 -sticky nwes
	frame := ui.NewFrame(root).
		SetPadding(3, 3, 12, 12)

	ui.StartGridConfig(frame).
		Sticky("nwes").
		Col(0).
		Row(0).
		Exec()

	// grid columnconfigure . 0 -weight 1; grid rowconfigure . 0 -weight 1
	ui.StartGridColConfig(root, 0).Weight(1).Exec()
	ui.StartGridRowConfig(root, 0).Weight(1).Exec()

	// grid [ttk::entry .c.feet -width 7 -textvariable feet] -column 2 -row 1 -sticky we
	feet := ui.NewEntry(frame).SetWidth(7)
	ui.StartGridConfig(feet).Col(2).Row(1).Sticky("we").Exec()

	// grid [ttk::label .c.meters -textvariable meters] -column 2 -row 2 -sticky we
	meters := ui.NewLabel(frame, "")
	ui.StartGridConfig(meters).Col(2).Row(2).Sticky("we").Exec()

	// grid [ttk::button .c.calc -text "Calculate" -command calculate] -column 3 -row 3 -sticky w
	// proc calculate {} {
	//   if {[catch {
	//       set ::meters [expr {round($::feet*0.3048*10000.0)/10000.0}]
	//   }]!=0} {
	//       set ::meters ""
	//   }
	// }

	calcChan := make(gotk.EventChannel)

	go func() {
		for e := range calcChan {
			fmt.Println(e)
			ftVal, err := strconv.Atoi(strings.TrimSpace(feet.Value()))
			if err != nil {
				meters.SetText(err.Error())
				return
			}

			rsltVal := math.Round(float64(ftVal) * 0.3048 * 10000.0) / 10000.0

			meters.SetText(fmt.Sprintf("%v", rsltVal))
		}
	}()

	calc := ui.NewButton(frame, "Calculate", calcChan)

	ui.StartGridConfig(calc).Col(3).Row(3).Sticky("w").Exec()

	// grid [ttk::label .c.flbl -text "feet"] -column 3 -row 1 -sticky w
	ui.StartGridConfig(ui.NewLabel(frame, "feet")).Col(3).Row(1).Sticky("w").Exec()
	// grid [ttk::label .c.islbl -text "is equivalent to"] -column 1 -row 2 -sticky e
	ui.StartGridConfig(ui.NewLabel(frame, "is equivalent to")).Col(1).Row(2).Sticky("e").Exec()
	// grid [ttk::label .c.mlbl -text "meters"] -column 3 -row 2 -sticky w
	ui.StartGridConfig(ui.NewLabel(frame, "meters")).Col(3).Row(2).Sticky("w").Exec()

	// foreach w [winfo children .c] {grid configure $w -padx 5 -pady 5}
	for _, child := range frame.Children() {
		ui.StartGridConfig(child).PadX(5).PadY(5).Exec()
	}

	// focus .c.feet
	ui.SetFocus(feet)
	// bind . <Return> {calculate}
	ui.SetBindKey(root, 0, gotk.Return, calcChan)

	// get geometry >> ui.WmGeometry(root)
	ui.WmSetGeometry(root, 350, 200, 100, 100)

	mousewheel := make(chan int, 50)

	ui.BindMouseWheel(root, 0, mousewheel)

	go func() {
		for {
			_ = <-mousewheel
		}
	}()

	rslt, _ := ui.FileOpenStart().
		SetMultiple(true).
		SetTitle("whatever").
		Exec()

	fmt.Println(rslt)

	ui.Wait()

//	ui.Close()
}


/*

    wm geometry . {2000x600+2038+-100} << leave {...} off for query
    . cget -class



*/
