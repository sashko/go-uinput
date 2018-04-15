package main

import "github.com/sashko/go-uinput"
import "time"

func main() {
	keyboard, err := uinput.CreateKeyboard()
	if err != nil {
		return
	}
	defer keyboard.Close()

	time.Sleep(time.Second * 1)

	// Press left Shift key, press G, release Shift key
	keyboard.KeyDown(uinput.KeyLeftShift)
	keyboard.KeyPress(uinput.KeyG)
	keyboard.KeyUp(uinput.KeyLeftShift)

	// Press O key
	keyboard.KeyPress(uinput.KeyO)
}
