package uinput

import (
	"fmt"
	"io"
	"os"
	"time"
)

// TouchScreen interface
type TouchScreen interface {
	Touch(x int32, y int32) error

	io.Closer
}

type vTouchScreen struct {
	devFile *os.File
}

func setupTouchScreen(devFile *os.File, minX int32, maxX int32, minY int32, maxY int32) error {
	var uinp uinputUserDev

	uinp.Name = uinputSetupNameToBytes([]byte("GoUinputDevice"))
	uinp.ID.BusType = BusVirtual
	uinp.ID.Vendor = 1
	uinp.ID.Product = 2
	uinp.ID.Version = 3

	uinp.AbsMin[AbsMtPositionX] = minX
	uinp.AbsMax[AbsMtPositionX] = maxX
	uinp.AbsMin[AbsMtPositionY] = minY
	uinp.AbsMax[AbsMtPositionY] = maxY

	buf, err := uinputUserDevToBuffer(uinp)
	if err != nil {
		goto err
	}

	err = ioctl(devFile, uiSetEvBit, uintptr(EvKey))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetEvBit, uintptr(EvAbs))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetKeyBit, uintptr(BtnTouch))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetAbsBit, uintptr(AbsMtPositionX))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_ABSBIT ioctl for X axis: %v", err)
		goto err
	}

	err = ioctl(devFile, uiSetAbsBit, uintptr(AbsMtPositionY))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_SET_ABSBIT ioctl for Y axis: %v", err)
		goto err
	}

	_, err = devFile.Write(buf)
	if err != nil {
		err = fmt.Errorf("Could not write uinputUserDev to device: %v", err)
		goto err
	}

	err = ioctl(devFile, uiDevCreate, uintptr(0))
	if err != nil {
		err = fmt.Errorf("Could not perform UI_DEV_CREATE ioctl: %v", err)
		goto err
	}

	time.Sleep(time.Millisecond * 200)

	return nil

err:
	destroyDevice(devFile)

	return err
}

// CreateTouchScreen creates virtual input device that emulates touch screen
func CreateTouchScreen(minX int32, maxX int32, minY int32, maxY int32) (TouchScreen, error) {
	dev, err := openUinputDev()
	if err != nil {
		return nil, err
	}

	err = setupTouchScreen(dev, minX, maxX, minY, maxY)
	if err != nil {
		return nil, err
	}

	return vTouchScreen{devFile: dev}, err
}

// Touch emits touch event
func (vts vTouchScreen) Touch(x int32, y int32) error {
	err := emitEvent(vts.devFile, EvAbs, AbsMtPositionX, x)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vts.devFile, EvAbs, AbsMtPositionY, y)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vts.devFile, EvSyn, 2, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vts.devFile, EvSyn, 0, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vts.devFile, EvSyn, 2, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	err = emitEvent(vts.devFile, EvSyn, 0, 0)
	if err != nil {
		return fmt.Errorf("emitEvent: %v", err)
	}

	return nil
}

func (vts vTouchScreen) Close() error {
	return destroyDevice(vts.devFile)
}
