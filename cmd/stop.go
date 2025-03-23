/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Terminates Background Process",
	Long:  `Terminates background process and stops the tool`,
	Run: func(cmd *cobra.Command, args []string) {

		executablePath, err := os.Executable()
		if err != nil {
			fmt.Println("❌ Error getting executable path:", err)
			return
		}
		execDir := filepath.Dir(executablePath)
		configDir := filepath.Join(execDir, "config")
		pidFile := filepath.Join(configDir, "ghs.pid")

		// Check if PID file exists
		if !utils.Exists(pidFile) {
			fmt.Println("❌ No tracking process found")
			return
		}

		// Read the PID file
		pidData, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Println("❌ Error reading process ID:", err)
			return
		}

		// Parse the PID
		pid, err := strconv.Atoi(string(pidData))
		if err != nil {
			fmt.Println("❌ Invalid process ID in file")
			os.Remove(pidFile)
			return
		}

		// Kill the process
		if err := killProcess(pid); err != nil {
			fmt.Println("❌ Failed to stop tracking process:", err)
			return
		}

		// Remove the PID file
		os.Remove(pidFile)
		fmt.Println("✅ Tracking stopped successfully")
	},
}

func killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process not found: %v", err)
	}

	if runtime.GOOS == "windows" {
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
	} else {
		if err := process.Signal(os.Interrupt); err != nil {
			if err := process.Kill(); err != nil {
				return fmt.Errorf("failed to kill process: %v", err)
			}
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
