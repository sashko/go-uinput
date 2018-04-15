package uinput

import (
	"fmt"
	"os"
	"syscall"
)

func openUinputDev() (devFile *os.File, err error) {
	devFile, err = os.OpenFile(uinputDevPath, os.O_WRONLY|syscall.O_NONBLOCK, 0660)
	if err != nil {
		return nil, fmt.Errorf("Could not open /dev/uinput: %v", err)
	}

	return devFile, err
}

func destroyDevice(devFile *os.File) error {
	err := ioctl(devFile, uiDevDestroy, uintptr(0))
	if err != nil {
		return fmt.Errorf("Could not perform UI_DEV_DESTOY ioctl: %v", err)
	}

	return devFile.Close()
}
