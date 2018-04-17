package main

import "github.com/sashko/go-uinput"
import "time"

func miceExample() {
	mice, err := uinput.CreateMice(0, 1919, 0, 1079)
	if err != nil {
		return
	}
	defer mice.Close()

	// draw a 500x500px square
	for i := 0; i <= 50; i++ {
		mice.MoveX(int32(10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mice.MoveY(int32(10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mice.MoveX(int32(-10))
		time.Sleep(time.Millisecond * 20)
	}
	for i := 0; i <= 50; i++ {
		mice.MoveY(int32(-10))
		time.Sleep(time.Millisecond * 20)
	}
}
