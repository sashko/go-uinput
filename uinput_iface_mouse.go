package uinput

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Mouse interface
type Mouse interface {
	LeftPress() error

	LeftRelease() error

	LeftClick() error

	RightPress() error

	RightRelease() error

	RightClick() error

	MiddleClick() error

	SideClick() error

	ExtraClick() error

	ForwardClick() error

	BackClick() error

	MoveX(x int32) error

	MoveY(x int32) error

	io.Closer
}

// Mice is the former name of Mouse.
//
// Deprecated: use Mouse.
type Mice = Mouse

type vMouse struct {
	devFile *os.File
}

func setupMouse(devFile *os.File) error {
	var uinp uinputUserDev

	uinp.Name = uinputSetupNameToBytes([]byte("GoUinputDevice"))
	uinp.ID.BusType = BusVirtual
	uinp.ID.Vendor = 1
	uinp.ID.Product = 2
	uinp.ID.Version = 3

	buf, err := uinputUserDevToBuffer(uinp)
	if err != nil {
		goto err
	}

	// register left and right buttons click events
	err = ioctl(devFile, uiSetEvBit, uintptr(EvKey))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnLeft))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnRight))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnMiddle))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnSide))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnExtra))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnForward))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnBack))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	// setup relative axes
	err = ioctl(devFile, uiSetEvBit, uintptr(EvRel))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetRelBit, uintptr(RelX))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_RELBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetRelBit, uintptr(RelY))
	if err != nil {
		err = fmt.Errorf("could not perform UI_SET_RELBIT ioctl: %v", err)
		goto err
	}

	_, err = devFile.Write(buf)
	if err != nil {
		err = fmt.Errorf("could not write uinputUserDev to device: %v", err)
		goto err
	}

	err = ioctl(devFile, uiDevCreate, uintptr(0))
	if err != nil {
		devFile.Close()
		return fmt.Errorf("could not perform UI_DEV_CREATE ioctl: %v", err)
	}

	time.Sleep(time.Millisecond * 200)

	return nil

err:
	destroyDevice(devFile)

	return err
}

// CreateMouse creates virtual input device that emulates mouse
func CreateMouse() (Mouse, error) {
	dev, err := openUinputDev()
	if err != nil {
		return nil, err
	}

	err = setupMouse(dev)
	if err != nil {
		return nil, err
	}

	return vMouse{devFile: dev}, err
}

// CreateMice creates virtual input device that emulates mice
//
// Deprecated: the range arguments have no effect, because the mouse reports
// relative motion only. Use CreateMouse.
func CreateMice(minX int32, maxX int32, minY int32, maxY int32) (Mice, error) {
	return CreateMouse()
}

// LeftPress emits left button press event
func (vm vMouse) LeftPress() error {
	err := emitEvent(vm.devFile, EvKey, BtnLeft, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftRelease emits left button release event
func (vm vMouse) LeftRelease() error {
	err := emitEvent(vm.devFile, EvKey, BtnLeft, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftClick emits left button click event
func (vm vMouse) LeftClick() error {
	err := vm.LeftPress()
	if err != nil {
		return err
	}

	return vm.LeftRelease()
}

// RightPress emits right button press event
func (vm vMouse) RightPress() error {
	err := emitEvent(vm.devFile, EvKey, BtnRight, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightRelease emits right button release event
func (vm vMouse) RightRelease() error {
	err := emitEvent(vm.devFile, EvKey, BtnRight, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightClick emits right button click event
func (vm vMouse) RightClick() error {
	err := vm.RightPress()
	if err != nil {
		return err
	}

	return vm.RightRelease()
}

// MiddleClick emits middle button click event
func (vm vMouse) MiddleClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnMiddle, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnMiddle, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// SideClick emits side button click event
func (vm vMouse) SideClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnSide, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnSide, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// ExtraClick emits extra button click event
func (vm vMouse) ExtraClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnExtra, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnExtra, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// ForwardClick emits forward button click event
func (vm vMouse) ForwardClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnForward, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnForward, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// BackClick emits back button click event
func (vm vMouse) BackClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnBack, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnBack, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// MoveX emits X axis movement event
func (vm vMouse) MoveX(x int32) error {
	err := emitEvent(vm.devFile, EvRel, RelX, x)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// MoveY emits Y axis movement event
func (vm vMouse) MoveY(x int32) error {
	err := emitEvent(vm.devFile, EvRel, RelY, x)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// Close destroys the virtual input device
func (vm vMouse) Close() error {
	return destroyDevice(vm.devFile)
}
