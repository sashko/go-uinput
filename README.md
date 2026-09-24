# go-uinput

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Coverage Status](https://coveralls.io/repos/github/sashko/go-uinput/badge.svg?branch=master)](https://coveralls.io/github/sashko/go-uinput?branch=master)

go-uinput is a Go interface to the Linux uinput kernel module that makes it possible to emulate input devices from userspace.

The interface makes it easy to create virtual input devices, such as keyboards, joysticks, or mice, to generate arbitrary input events programmatically.

## System prerequisites

First, the system must have the `uinput` kernel module loaded

    sudo modprobe -i uinput

Second, the `/dev/uinput` device is owned by root, and therefore its default permissions must either be changed using chmod

    sudo chmod 666 /dev/uinput

or, which is much preferred, add the udev rule to allow a user to use the device

    echo KERNEL=="uinput", MODE="0666" | sudo tee /etc/udev/rules.d/90-$USER.rules
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
