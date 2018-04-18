package uinput

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Mice interface
type Mice interface {
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

type vMice struct {
	devFile *os.File
}

func setupMice(devFile *os.File, minX int32, maxX int32, minY int32, maxY int32) error {
	var uinp uinputUserDev

	uinp.Name = uinputSetupNameToBytes([]byte("GoUinputDevice"))
	uinp.ID.BusType = BusVirtual
	uinp.ID.Vendor = 1
	uinp.ID.Product = 2
	uinp.ID.Version = 3

	uinp.AbsMin[AbsX] = minX
	uinp.AbsMax[AbsX] = maxX
	uinp.AbsFuzz[AbsX] = 0
	uinp.AbsFlat[AbsX] = 0
	uinp.AbsMin[AbsY] = minY
	uinp.AbsMax[AbsY] = maxY
	uinp.AbsFuzz[AbsY] = 0
	uinp.AbsFlat[AbsY] = 0

	buf, err := uinputUserDevToBuffer(uinp)
	if err != nil {
		goto err
	}

	// register left and right buttons click events
	err = ioctl(devFile, uiSetEvBit, uintptr(EvKey))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnLeft))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnRight))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnMiddle))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnSide))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnExtra))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnForward))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnBack))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	// setup relative axes
	err = ioctl(devFile, uiSetEvBit, uintptr(EvRel))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetRelBit, uintptr(RelX))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetRelBit, uintptr(RelY))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	_, err = devFile.Write(buf)
	if err != nil {
		err = fmt.Errorf("Could not write uinputUserDev to device: %v", err)
		goto err
	}

	err = ioctl(devFile, uiDevCreate, uintptr(0))
	if err != nil {
		devFile.Close()
		return fmt.Errorf("Could not perform UI_DEV_CREATE ioctl: %v", err)
	}

	time.Sleep(time.Millisecond * 1000)

	return nil

err:
	destroyDevice(devFile)

	return err
}

// CreateMice creates virtual input device that emulates mice
func CreateMice(minX int32, maxX int32, minY int32, maxY int32) (Mice, error) {
	dev, err := openUinputDev()
	if err != nil {
		return nil, err
	}

	err = setupMice(dev, minX, maxX, minY, maxY)
	if err != nil {
		return nil, err
	}

	return vMice{devFile: dev}, err
}

// LeftPress emits left button press event
func (vm vMice) LeftPress() error {
	err := emitEvent(vm.devFile, EvKey, BtnLeft, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftRelease emits left button release event
func (vm vMice) LeftRelease() error {
	err := emitEvent(vm.devFile, EvKey, BtnLeft, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftClick emits left button click event
func (vm vMice) LeftClick() error {
	err := vm.LeftPress()
	err = vm.LeftRelease()

	return err
}

// RightPress emits right button press event
func (vm vMice) RightPress() error {
	err := emitEvent(vm.devFile, EvKey, BtnRight, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightPress emits right button release event
func (vm vMice) RightRelease() error {
	err := emitEvent(vm.devFile, EvKey, BtnRight, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightClick emits right button click event
func (vm vMice) RightClick() error {
	err := vm.RightPress()
	err = vm.RightRelease()

	return err
}

// MiddleClick emits middle button click event
func (vm vMice) MiddleClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnMiddle, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnMiddle, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// SideClick emits side button click event
func (vm vMice) SideClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnSide, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnSide, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// ExtraClick emits extra button click event
func (vm vMice) ExtraClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnExtra, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnExtra, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// ForwardClick emits forward button click event
func (vm vMice) ForwardClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnForward, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnForward, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// BackClick emits back button click event
func (vm vMice) BackClick() error {
	err := emitEvent(vm.devFile, EvKey, BtnBack, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvKey, BtnBack, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// MoveX emits X axis movement event
func (vm vMice) MoveX(x int32) error {
	err := emitEvent(vm.devFile, EvRel, RelX, x)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// MoveY emits Y axis movement event
func (vm vMice) MoveY(x int32) error {
	err := emitEvent(vm.devFile, EvRel, RelY, x)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vm.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

func (vm vMice) Close() error {
	return destroyDevice(vm.devFile)
}
