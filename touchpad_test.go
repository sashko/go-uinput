package uinput

import (
	"testing"
)

func TestTouchPadCreation(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.Close()
	if err != nil {
		t.Fatal("Failed to close virtual touchpad device")
	}
}

func TestVirtualTouchPadLeftPressAndRelease(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.LeftPress()
	if err != nil {
		t.Fatal("Failed to emit left button press")
	}

	err = touchPad.LeftRelease()
	if err != nil {
		t.Fatal("Failed to emit left button release")
	}
}

func TestVirtualTouchPadRightPressAndRelease(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.RightPress()
	if err != nil {
		t.Fatal("Failed to emit right button press")
	}

	err = touchPad.RightRelease()
	if err != nil {
		t.Fatal("Failed to emit right button release")
	}
}

func TestVirtualTouchPadLeftClick(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.LeftClick()
	if err != nil {
		t.Fatal("Failed to emit left button click")
	}
}

func TestVirtualTouchPadRightClick(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.RightClick()
	if err != nil {
		t.Fatal("Failed to emit right button click")
	}
}

func TestVirtualTouchPadMoveTo(t *testing.T) {
	touchPad, err := CreateTouchPad(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual touchpad")
	}

	err = touchPad.MoveTo(100, 200)
	if err != nil {
		t.Fatal("Failed to emit move to event")
	}
}
