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

func uinputSetupNameToBytes(name []byte) (uinputName [uinputMaxNameSize]byte) {
	var bytesName [uinputMaxNameSize]byte

	copy(bytesName[:], name)

	return bytesName
}

func inputEventToBuffer(iev inputEvent) (buffer []byte, err error) {
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.LittleEndian, iev)
	if err != nil {
		return nil, fmt.Errorf("Failed to write inputEvent to buffer: %v", err)
	}

	return buf.Bytes(), nil
}
