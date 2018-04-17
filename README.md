# go-uinput

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/sashko/go-uinput.svg?branch=master)](https://travis-ci.org/sashko/go-uinput)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashko/go-uinput)](https://goreportcard.com/report/github.com/sashko/go-uinput)

go-uinput is Go interface to Linux uinput kernel module that makes it possible to emulate input devices from userspace.

The interface aims to make it dead simple to create virtual input devices, e.g. keyboard, joystick, or mice for generating arbitrary input events programmatically.

## System prerequisites

First, the system must have `uinput` kernel module loaded.

    sudo modprobe -i uinput

Second, `/dev/uinput` device is owned by root and therefore its default permissions must either be changed using chmod

    sudo chmod 666 /dev/uinput

or, which is a much more preferred option, add the udev rule to allow a user use the device

    echo KERNEL=="uinput", MODE="0666" | sudo tee /etc/udev/rules.d/90-$USER.rules
    sudo udevadm trigger

## Installation

To build the package from sources, clone the repository and run

    go build
    go install

Alternatively, use `go get` to obtain and install the package in your `$GOPATH`

    go get github.com/sashko/go-uinput

## Usage

The following example shows how to create a new virtual keyboard, and how to send a key press event.

All the default imports and error handlers were removed for the sake of simplicity.

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
