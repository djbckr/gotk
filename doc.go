/*
	TK is a cross-platform UI scripting language. It is dynamic (not strongly typed) along
	the same vein as python or javascript. Go is a highly structured language that is great
	except that it does not have a UI component. Package gotk is  a bridge that allows
	a Go programmer to get a working, cross-platform UI into a Go program. It requires the use of
	the wish interpreter. If gotk can't find wish, it will fatally crash
	with a message that indicates that. See https://www.tcl.tk/ for more information.
	OSX already has TK/wish installed, Linux may need an installation, and Windows
	will need an installation.

	Below, we will outline some of the main features. Be sure to read the gotk_test.go file - most
	of the below examples come from that file. It can be found here: https://github.com/djbckr/gotk


	Usage

	You start a UI session by calling Tk()

		ui := gotk.Tk()

	This starts the wish interpreter and gives you a root window. You can conceivably have more than
	one wish interpreter running at the same time.

	  root := ui.Root()

	You can set the title of the root window:

		root.SetTitle("My Window Title")

	From here, you will typically set a frame on the root window:

	  frame := ui.NewFrame(root).
			SetPadding(3, 3, 12, 12)

	gotk implements most of the grid configuration functions, in which you call StartGridConfig(),
	then a series of functions to fully configure the grid, then you call .Exec() to finish and
	send the script to wish.

		ui.StartGridConfig(frame).
			Sticky("nwes").
			Col(0).
			Row(0).
			Exec()

	This method of calling multiple functions on a component allows gotk to build a dynamic script
	which is sent to wish. You will see this construct a lot.

	To interact with the mouse and keyboard, you will make channels that will have messages put in them
	when the appropriate event occurs. You then pass the channel into buttons or key bindings as appropriate.

		calcChan := make(gotk.EventChannel)

		go func() {
			for range calcChan {
				ftVal, err := strconv.Atoi(strings.TrimSpace(feet.Value()))
				if err != nil {
					meters.SetText(err.Error())
					return
				}

				rsltVal := math.Round(float64(ftVal) * 0.3048 * 10000.0) / 10000.0

				meters.SetText(fmt.Sprintf("%v", rsltVal))
			}
		}()

	Above, we have created a channel, and we setup a goroutine to listen to any incoming events
	on that channel. Then we pass that channel to a button or a bind key:

		calc := ui.NewButton(frame, "Calculate", calcChan)

		ui.SetBindKey(root, 0, gotk.Return, calcChan)

	Now, when the calc button is clicked, or when the Return key is pressed, an event will be placed
	into the calcChan channel, and your code will execute.

*/
package gotk
