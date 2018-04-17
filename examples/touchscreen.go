package main

import "github.com/sashko/go-uinput"
import "time"

func touchScreenExample() {
	touchScreen, err := uinput.CreateTouchScreen(0, 1079, 0, 719)
	if err != nil {
		return
	}
	defer touchScreen.Close()

	for i := 0; i <= 1079; i += 100 {
		for j := 0; j <= 719; j += 100 {
			touchScreen.Touch(int32(i), int32(j))
		}
	}
}
