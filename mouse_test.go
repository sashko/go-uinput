package uinput

import "testing"

func TestVirtualMouseCreation(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseDeprecatedCreation(t *testing.T) {
	mouse, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseLeftPressAndRelease(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	err = mouse.LeftPress()
	if err != nil {
		t.Fatal("Failed to emit left button press")
	}

	err = mouse.LeftRelease()
	if err != nil {
		t.Fatal("Failed to emit left button release")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseRightPressAndRelease(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	err = mouse.RightPress()
	if err != nil {
		t.Fatal("Failed to emit right button press")
	}

	err = mouse.RightRelease()
	if err != nil {
		t.Fatal("Failed to emit right button release")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseLeftClick(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	err = mouse.LeftClick()
	if err != nil {
		t.Fatal("Failed to emit left button click")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseRightClick(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	err = mouse.RightClick()
	if err != nil {
		t.Fatal("Failed to emit right button click")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseExtraButtonsClick(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	err = mouse.MiddleClick()
	if err != nil {
		t.Fatal("Failed to emit middle button click")
	}

	err = mouse.SideClick()
	if err != nil {
		t.Fatal("Failed to emit side button click")
	}

	err = mouse.ExtraClick()
	if err != nil {
		t.Fatal("Failed to emit extra button click")
	}

	err = mouse.ForwardClick()
	if err != nil {
		t.Fatal("Failed to emit forward button click")
	}

	err = mouse.BackClick()
	if err != nil {
		t.Fatal("Failed to emit back button click")
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}

func TestVirtualMouseXYAxisMovement(t *testing.T) {
	mouse, err := CreateMouse()
	if err != nil {
		t.Fatal("Failed to create virtual mouse")
	}

	grabDevice(t, mouse.(vMouse).devFile)

	for i := 0; i <= 50; i++ {
		err = mouse.MoveX(int32(10))
		if err != nil {
			t.Fatal("Failed to move cursor to the right along the x axis")
		}
	}

	for i := 0; i <= 50; i++ {
		err = mouse.MoveY(int32(10))
		if err != nil {
			t.Fatal("Failed to move cursor down along the Y axis")
		}
	}

	for i := 0; i <= 50; i++ {
		err = mouse.MoveX(int32(-10))
		if err != nil {
			t.Fatal("Failed to move cursor to the left along the x axis")
		}
	}

	for i := 0; i <= 50; i++ {
		err = mouse.MoveY(int32(-10))
		if err != nil {
			t.Fatal("Failed to move cursor up along the Y axis")
		}
	}

	err = mouse.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mouse device")
	}
}
