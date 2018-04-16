package uinput

import "testing"

func TestVirtualKeyboard(t *testing.T) {
	keyboard, err := CreateKeyboard()
	if err != nil {
		t.Fatal("Failed to create virtual keyboard")
	}

	for i := 0; i < KeyMax; i++ {
		err = keyboard.KeyPress(uint16(i))
		if err != nil {
			t.Fatal("Failed to press key")
		}
	}

	err = keyboard.Close()
	if err != nil {
		t.Fatal("Failed to close virtual keyboard device")
	}
}
