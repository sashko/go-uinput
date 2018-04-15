package uinput

import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

// Keyboard interface
type Keyboard interface {
	KeyPress(key uint16) error

	KeyDown(key uint16) error

	KeyUp(key uint16) error

	io.Closer
}

type vKeyboard struct {
	devFile *os.File
}

func setupKeyboard(devFile *os.File) error {
	err := ioctl(devFile, uiSetEvBit, EvKey)
	if err != nil {
		return fmt.Errorf("Could not perform UI_SET_EVBIT ioctl: %v", err)
	}

	for i := 0; i < KeyMax; i++ {
		err = ioctl(devFile, uiSetKeyBit, uintptr(i))
		if err != nil {
			return fmt.Errorf("Could not perform UI_SET_KEYBIT ioctl: %v", err)
		}
	}

	var usetup uinputSetup

	// TODO: add possibility to change those values
	usetup.name = uinputSetupNameToBytes([]byte("GoUinputDevice"))
	usetup.id.busType = BusUSB
	usetup.id.vendor = 1
	usetup.id.product = 2
	usetup.id.version = 3

	err = ioctl(devFile, uiDevSetup, uintptr(unsafe.Pointer(&usetup)))
	if err != nil {
		return fmt.Errorf("Could not perform UI_DEV_SETUP ioctl: %v", err)
	}

	err = ioctl(devFile, uiDevCreate, uintptr(0))
	if err != nil {
		return fmt.Errorf("Could not perform UI_DEV_CREATE ioctl: %v", err)
	}

	return err
}

// CreateKeyboard creates virtual keyboard
func CreateKeyboard() (Keyboard, error) {
	dev, err := openUinputDev()
	if err != nil {
		return nil, err
	}

	err = setupKeyboard(dev)
	if err != nil {
		return nil, err
	}

	return vKeyboard{devFile: dev}, err
}

func (vk vKeyboard) KeyPress(key uint16) error {
	err := emitKeyDown(vk.devFile, key)
	if err != nil {
		return err
	}

	err = emitKeyUp(vk.devFile, key)
	if err != nil {
		return err
	}

	return err
}

func (vk vKeyboard) KeyDown(key uint16) error {
	err := emitKeyDown(vk.devFile, key)
	if err != nil {
		return err
	}

	return err
}

func (vk vKeyboard) KeyUp(key uint16) error {
	err := emitKeyUp(vk.devFile, key)
	if err != nil {
		return err
	}

	return err
}

func (vk vKeyboard) Close() error {
	return destroyDevice(vk.devFile)
}

func emitKeyEvent(devFile *os.File, typ uint16, code uint16, value int32) error {
	var ie inputEvent

	ie.Type = typ
	ie.Code = code
	ie.Value = value

	buf, err := inputEventToBuffer(ie)
	if err != nil {
		return fmt.Errorf("Could not write inputEvent to buffer: %v", err)
	}

	_, err = devFile.Write(buf)
	if err != nil {
		return fmt.Errorf("Could write to the device: %v", err)
	}

	return nil
}

func emitKeyDown(devFile *os.File, code uint16) error {
	err := emitKeyEvent(devFile, EvKey, code, 1)
	if err != nil {
		return fmt.Errorf("Could not emit key down event: %v", err)
	}

	err = emitKeyEvent(devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("Could not emit sync event: %v", err)
	}

	return err
}

func emitKeyUp(devFile *os.File, code uint16) error {
	err := emitKeyEvent(devFile, EvKey, code, 0)
	if err != nil {
		return fmt.Errorf("Could not emit key up event: %v", err)
	}

	err = emitKeyEvent(devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("Could not emit sync event: %v", err)
	}

	return err
}
