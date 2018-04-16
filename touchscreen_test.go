package uinput

import (
	"testing"
	"time"
)

func TestVirtualTouchScreen(t *testing.T) {
	touchScreen, err := CreateTouchScreen(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchScreen")
	}

	for i := 0; i < 1000; i += 100 {
		for j := 0; j < 700; j += 100 {
			err = touchScreen.Touch(int32(i), int32(j))
			if err != nil {
				t.Fatal("Failed to touch screen")
			}

			time.Sleep(time.Millisecond * 200)
		}
	}

	err = touchScreen.Close()
	if err != nil {
		t.Fatal("Failed to close virtual touchScreen device")
	}
}
