package uinput

import (
	"testing"
	"time"
)

func TestVirtualMiceCreation(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.Close()
	if err != nil {
		t.Fatal("Failed to close virtual mice device")
	}
}

func TestVirtualMiceLeftPressAndRelease(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.LeftPress()
	if err != nil {
		t.Fatal("Failed to emit left button press")
	}

	err = mice.LeftRelease()
	if err != nil {
		t.Fatal("Failed to emit left button release")
	}
}

func TestVirtualMiceRightPressAndRelease(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.RightPress()
	if err != nil {
		t.Fatal("Failed to emit right button press")
	}

	err = mice.RightRelease()
	if err != nil {
		t.Fatal("Failed to emit right button release")
	}
}

func TestVirtualMiceLeftClick(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.LeftClick()
	if err != nil {
		t.Fatal("Failed to emit left button click")
	}
}

func TestVirtualMiceRightClick(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.RightClick()
	if err != nil {
		t.Fatal("Failed to emit right button click")
	}
}

func TestVirtualMiceExtraButtonsClick(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	err = mice.MiddleClick()
	if err != nil {
		t.Fatal("Failed to emit middle button click")
	}

	err = mice.SideClick()
	if err != nil {
		t.Fatal("Failed to emit side button click")
	}

	err = mice.ExtraClick()
	if err != nil {
		t.Fatal("Failed to emit extra button click")
	}

	err = mice.ForwardClick()
	if err != nil {
		t.Fatal("Failed to emit forward button click")
	}

	err = mice.BackClick()
	if err != nil {
		t.Fatal("Failed to emit middle button click")
	}
}

func TestVirtualMiceXYAxisMovement(t *testing.T) {
	mice, err := CreateMice(0, 1079, 0, 719)
	if err != nil {
		t.Fatal("Failed to create virtual mice")
	}

	for i := 0; i <= 50; i++ {
		err = mice.MoveX(int32(10))
		if err != nil {
			t.Fatal("Failed to move cursor to the right along the x axis")
		}
		time.Sleep(time.Millisecond * 5)
	}

	for i := 0; i <= 50; i++ {
		mice.MoveY(int32(10))
		if err != nil {
			t.Fatal("Failed to move cursor down along the Y axis")
		}
		time.Sleep(time.Millisecond * 20)
	}

	for i := 0; i <= 50; i++ {
		err = mice.MoveX(int32(-10))
		if err != nil {
			t.Fatal("Failed to move cursor to the left along the x axis")
		}
		time.Sleep(time.Millisecond * 20)
	}

	for i := 0; i <= 50; i++ {
		err = mice.MoveY(int32(-10))
		if err != nil {
			t.Fatal("Failed to move cursor up along the Y axis")
		}
		time.Sleep(time.Millisecond * 20)
	}
}
