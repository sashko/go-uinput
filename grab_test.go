package uinput

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
	"unsafe"
)

const (
	uiGetSysname = 0x8040552c // UI_GET_SYSNAME(64) from uinput.h
	evIocGrab    = 0x40044590 // EVIOCGRAB from input.h

	grabTimeout = time.Second
)

// grabDevice takes exclusive access to the events of the virtual device
// behind devFile, so they never reach the desktop or systemd-logind (which
// acts on keys such as KEY_POWER and KEY_RESTART). It must be called before
// any event is emitted. The test fails, without emitting anything, if the
// device cannot be grabbed.
func grabDevice(t *testing.T, devFile *os.File) {
	t.Helper()

	var sysname [64]byte

	err := ioctlPtr(devFile, uiGetSysname, unsafe.Pointer(&sysname[0]))
	if err != nil {
		t.Fatalf("Failed to get sysname of virtual device: %v", err)
	}

	sysDir := filepath.Join("/sys/devices/virtual/input", string(bytes.TrimRight(sysname[:], "\x00")))

	// udev may still be creating the event node or applying its permissions.
	var evFile *os.File

	deadline := time.Now().Add(grabTimeout)
	for {
		nodes, _ := filepath.Glob(filepath.Join(sysDir, "event*"))
		if len(nodes) != 1 {
			err = os.ErrNotExist
		} else {
			evFile, err = os.OpenFile(filepath.Join("/dev/input", filepath.Base(nodes[0])), os.O_RDONLY, 0)
			if err == nil {
				break
			}
		}

		if time.Now().After(deadline) {
			t.Fatalf("Failed to open event node of %s (read access to /dev/input/event* is required): %v", sysDir, err)
		}

		time.Sleep(time.Millisecond * 20)
	}

	err = ioctl(evFile, evIocGrab, 1)
	if err != nil {
		evFile.Close()
		t.Fatalf("Failed to grab virtual device: %v", err)
	}

	t.Cleanup(func() { evFile.Close() })
}
