package uinput_test

import (
	"log"
	"time"

	"github.com/sashko/go-uinput"
)

// The examples have no "Output:" comment, so go test compiles them but never
// runs them. Running them by hand sends real input to the desktop.

func ExampleCreateKeyboard() {
	keyboard, err := uinput.CreateKeyboard()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Press left Shift key, press G, release Shift key
	err = keyboard.KeyDown(uinput.KeyLeftShift)
	if err != nil {
		log.Fatal(err)
	}

	err = keyboard.KeyPress(uinput.KeyG)
	if err != nil {
		log.Fatal(err)
	}

	err = keyboard.KeyUp(uinput.KeyLeftShift)
	if err != nil {
		log.Fatal(err)
	}

	// Press O key
	err = keyboard.KeyPress(uinput.KeyO)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleCreateMouse() {
	mouse, err := uinput.CreateMouse()
	if err != nil {
		log.Fatal(err)
	}
	defer mouse.Close()

	// Draw a 500x500 px square in 50 steps of 10 px per side
	moves := []struct {
		move func(int32) error
		step int32
	}{
		{mouse.MoveX, 10},
		{mouse.MoveY, 10},
		{mouse.MoveX, -10},
		{mouse.MoveY, -10},
	}

	for _, m := range moves {
		for i := 0; i < 50; i++ {
			err = m.move(m.step)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Millisecond * 20)
		}
	}
}

func ExampleCreateTouchPad() {
	touchPad, err := uinput.CreateTouchPad(0, 1919, 0, 1079)
	if err != nil {
		log.Fatal(err)
	}
	defer touchPad.Close()

	err = touchPad.MoveTo(300, 200)
	if err != nil {
		log.Fatal(err)
	}

	err = touchPad.RightClick()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleCreateTouchScreen() {
	touchScreen, err := uinput.CreateTouchScreen(0, 1919, 0, 1079)
	if err != nil {
		log.Fatal(err)
	}
	defer touchScreen.Close()

	// Tap a grid of points 200 px apart
	for x := int32(0); x <= 1919; x += 200 {
		for y := int32(0); y <= 1079; y += 200 {
			err = touchScreen.Touch(x, y)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
