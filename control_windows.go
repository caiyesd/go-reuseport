package reuseport

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func GetBindToDeviceControl(device string) func(network, address string, c syscall.RawConn) error {
	return Control
}

func Control(network, address string, c syscall.RawConn) (err error) {
	return c.Control(func(fd uintptr) {
		err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
	})
}
