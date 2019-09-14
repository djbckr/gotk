# GO-TK
This is a go (golang) library that allows you to have a
cross-platform UI using the Tcl/Tk library. Rather than interfacing
directly to a C library, this uses the `wish` program when TK
is installed on your computer. OSX already has it installed; for Windows
and Linux, you'll need to install it. Information can be found
[here](https://www.tcl.tk/software/tcltk/)

#### How it works
As noted, the interface between go and Tcl/Tk is through the `wish`
program. Your Go program sends commands to `wish`, and when you want
information, such as the contents of an entry field, or a reaction
to a button click, `wish` sends that information back to your Go program
via network sockets. As a result, there is no messing about with
`unsafe` memory management calling into C libraries.

#### Limitations
At this point, the library is strictly UI-centric. It does not support
any of the Tcl commands, though you can send raw commands using
this library if you want. The intent is to have Go create a UI, and
the events (button clicks, primarily) will call Go functions.

For a quick example, see the gotk_test.go file.
This is a work in progress, not ready for prime-time.
