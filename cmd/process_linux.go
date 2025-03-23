//go:build linux
// +build linux

package cmd

import (
	"fmt"
	"os"
)

func killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process not found: %v", err)
	}

	if err := process.Signal(os.Interrupt); err != nil {
		if err := process.Kill(); err != nil {
			return fmt.Errorf("failed to kill process: %v", err)
		}
	}

	return nil
}
