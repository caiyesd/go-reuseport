package reuseport

import (
	"syscall"
)

func GetBindToDeviceControl(device string) func(network, address string, c syscall.RawConn) error {
	return Control
}

func Control(network, address string, c syscall.RawConn) error {
	return nil
}
