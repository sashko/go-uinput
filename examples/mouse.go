package main

import "github.com/sashko/go-uinput"
import "time"

func mouseExample() {
	mouse, err := uinput.CreateMouse()
	if err != nil {
		return
	}
	defer mouse.Close()

	// draw a 500x500px square
	for i := 0; i <= 50; i++ {
		mouse.MoveX(int32(10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mouse.MoveY(int32(10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mouse.MoveX(int32(-10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mouse.MoveY(int32(-10))
		time.Sleep(time.Millisecond * 20)
	}
}
