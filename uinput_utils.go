package uinput

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
)

func ioctl(df *os.File, op, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, df.Fd(), op, arg)
	if err != 0 {
		return syscall.Errno(err)
	}

	return nil
}

func emitEvent(devFile *os.File, typ uint16, code uint16, value int32) error {
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
		return fmt.Errorf("Could write inputEvent to device: %v", err)
	}

	return nil
}

func uinputSetupNameToBytes(name []byte) (uinputName [uinputMaxNameSize]byte) {
	var bytesName [uinputMaxNameSize]byte

	copy(bytesName[:], name)

	return bytesName
}

func inputEventToBuffer(iev inputEvent) (buffer []byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, iev)
	if err != nil {
		return nil, fmt.Errorf("Could not write inputEvent to buffer: %v", err)
	}

	return buf.Bytes(), nil
}

func uinputUserDevToBuffer(uud uinputUserDev) (buffer []byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, uud)
	if err != nil {
		return nil, fmt.Errorf("Could not write uinputUserDev to buffer: %v", err)
	}

	return buf.Bytes(), nil
}
