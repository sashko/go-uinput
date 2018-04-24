package uinput

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TouchPad interface {
	LeftPress() error

	LeftRelease() error

	LeftClick() error

	RightPress() error

	RightRelease() error

	RightClick() error

	MoveTo(x int32, y int32) error

	io.Closer
}

type vTouchPad struct {
	devFile *os.File
}

func setupTouchPad(devFile *os.File, minX int32, maxX int32, minY int32, maxY int32) error {
	var uinp uinputUserDev

	uinp.Name = uinputSetupNameToBytes([]byte("GoUinputDevice"))
	uinp.ID.BusType = BusVirtual
	uinp.ID.Vendor = 1
	uinp.ID.Product = 2
	uinp.ID.Version = 3

	uinp.AbsMin[AbsX] = minX
	uinp.AbsMax[AbsX] = maxX
	uinp.AbsMin[AbsY] = minY
	uinp.AbsMax[AbsY] = maxY

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

	// setup absolute axes
	err = ioctl(devFile, uiSetEvBit, uintptr(EvAbs))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetAbsBit, uintptr(AbsX))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_ABSBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetAbsBit, uintptr(AbsY))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_ABSBIT ioctl: %v", err)
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

	time.Sleep(time.Millisecond * 200)

	return nil

err:
	destroyDevice(devFile)

	return err
}

func CreateTouchPad(minX int32, maxX int32, minY int32, maxY int32) (TouchPad, error) {
	dev, err := openUinputDev()
	if err != nil {
		return nil, err
	}

	err = setupTouchPad(dev, minX, maxX, minY, maxY)
	if err != nil {
		return nil, err
	}

	return vTouchPad{devFile: dev}, err
}

// LeftPress emits left button press event
func (tp vTouchPad) LeftPress() error {
	err := emitEvent(tp.devFile, EvKey, BtnLeft, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(tp.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftRelease emits left button release event
func (tp vTouchPad) LeftRelease() error {
	err := emitEvent(tp.devFile, EvKey, BtnLeft, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(tp.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// LeftClick emits left button click event
func (tp vTouchPad) LeftClick() error {
	err := tp.LeftPress()
	err = tp.LeftRelease()

	return err
}

// RightPress emits right button press event
func (tp vTouchPad) RightPress() error {
	err := emitEvent(tp.devFile, EvKey, BtnRight, 1)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(tp.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightPress emits right button release event
func (tp vTouchPad) RightRelease() error {
	err := emitEvent(tp.devFile, EvKey, BtnRight, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(tp.devFile, EvSyn, 0, SynReport)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

// RightClick emits right button click event
func (tp vTouchPad) RightClick() error {
	err := tp.RightPress()
	err = tp.RightRelease()

	return err
}

func (tp vTouchPad) MoveTo(x int32, y int32) error {
	err := emitAbsEvent(tp.devFile, x, y)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

func (tp vTouchPad) Close() error {
	return destroyDevice(tp.devFile)
}
