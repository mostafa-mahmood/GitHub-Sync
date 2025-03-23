//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"syscall"
)

func killProcess(pid int) error {
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return fmt.Errorf("failed to load kernel32.dll: %v", err)
	}
	defer dll.Release()

	proc, err := dll.FindProc("TerminateProcess")
	if err != nil {
		return fmt.Errorf("failed to find TerminateProcess: %v", err)
	}

	handle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return fmt.Errorf("failed to open process: %v", err)
	}
	defer syscall.CloseHandle(handle)

	ret, _, err := proc.Call(uintptr(handle), 0)
	if ret == 0 {
		return fmt.Errorf("failed to terminate process: %v", err)
	}

	return nil
}
