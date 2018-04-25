package main

import "github.com/sashko/go-uinput"

func touchPadExample() {
	touchPad, err := uinput.CreateTouchPad(0, 1919, 0, 1079)
	if err != nil {
		return
	}
	defer touchPad.Close()

	touchPad.MoveTo(300, 200)

	touchPad.RightClick()
}
