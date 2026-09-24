# go-uinput

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Coverage Status](https://coveralls.io/repos/github/sashko/go-uinput/badge.svg?branch=master)](https://coveralls.io/github/sashko/go-uinput?branch=master)

go-uinput is a Go interface to the Linux uinput kernel module that makes it possible to emulate input devices from userspace.

The interface makes it easy to create virtual input devices, such as keyboards, joysticks, or mice, to generate arbitrary input events programmatically.

## System prerequisites

First, the system must have the `uinput` kernel module loaded

    sudo modprobe -i uinput

Second, the `/dev/uinput` device is owned by root, so a regular user needs access to it.

Note that any process with access to `/dev/uinput` can inject keystrokes and pointer events into every session on the system, including a root terminal. Grant access only to the users who need it, and never make the device world-writable (`chmod 666` or `MODE="0666"`).

The preferred way is a udev rule that grants access to the user logged in at the local seat

    echo 'KERNEL=="uinput", SUBSYSTEM=="misc", TAG+="uaccess", OPTIONS+="static_node=uinput"' | sudo tee /etc/udev/rules.d/60-uinput.rules
    sudo udevadm control --reload-rules
    sudo udevadm trigger

The rule file name must sort before `73-seat-late.rules`, which applies the `uaccess` tag.

For users without a local session, such as services or SSH logins, use a dedicated group instead, then log out and back in

    sudo groupadd --system uinput
    sudo usermod -aG uinput $USER
    echo 'KERNEL=="uinput", SUBSYSTEM=="misc", GROUP="uinput", MODE="0660", OPTIONS+="static_node=uinput"' | sudo tee /etc/udev/rules.d/60-uinput.rules
    sudo udevadm control --reload-rules
    sudo udevadm trigger

## Installation

go-uinput requires Go 1.22 or later. To add it to your module, run

    go get github.com/sashko/go-uinput@latest

and import it as

```go
import "github.com/sashko/go-uinput"
```

## Usage

The following example shows how to create a new virtual keyboard and how to send a key press event.

For simplicity, we removed all default imports and error handlers.

```go
func main() {
	keyboard, err := uinput.CreateKeyboard()

	defer keyboard.Close()

	// Press left Shift key, press G, release Shift
	keyboard.KeyDown(uinput.KeyLeftShift)
	keyboard.KeyPress(uinput.KeyG)
	keyboard.KeyUp(uinput.KeyLeftShift)

	// Press O key
	keyboard.KeyPress(uinput.KeyO)
}
```
