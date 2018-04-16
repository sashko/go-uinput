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
	usetup.id.BusType = BusUSB
	usetup.id.Vendor = 1
	usetup.id.Product = 2
	usetup.id.Version = 3

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

func emitKeyDown(devFile *os.File, code uint16) error {
	err := emitEvent(devFile, EvKey, code, 1)
	if err != nil {
		return fmt.Errorf("Could not emit key down event: %v", err)
	}

	err = emitEvent(devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("Could not emit sync event: %v", err)
	}

	return err
}

func emitKeyUp(devFile *os.File, code uint16) error {
	err := emitEvent(devFile, EvKey, code, 0)
	if err != nil {
		return fmt.Errorf("Could not emit key up event: %v", err)
	}

	err = emitEvent(devFile, EvSyn, SynReport, 0)
	if err != nil {
		return fmt.Errorf("Could not emit sync event: %v", err)
	}

	return err
}
