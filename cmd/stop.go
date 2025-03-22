/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Terminates Background Process",
	Long:  `Terminates background process and stop the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if PID file exists
		if !utils.Exists("./config/ghs.pid") {
			fmt.Println("❌ No tracking process found")
			return
		}

		// Read the PID file
		pidData, err := os.ReadFile("./config/ghs.pid")
		if err != nil {
			fmt.Println("❌ Error reading process ID:", err)
			return
		}

		// Parse the PID
		pid, err := strconv.Atoi(string(pidData))
		if err != nil {
			fmt.Println("❌ Invalid process ID in file")
			os.Remove("./config/ghs.pid")
			return
		}

		// Find the process
		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println("❌ Process not found")
			os.Remove("./config/ghs.pid")
			return
		}

		// Kill the process
		err = process.Signal(syscall.SIGTERM) // Try graceful termination first
		if err != nil {
			// If graceful termination fails, try force kill
			err = process.Kill()
			if err != nil {
				fmt.Println("❌ Failed to stop tracking process:", err)
				return
			}
		}

		// Remove the PID file
		os.Remove("./config/ghs.pid")
		fmt.Println("✅ Tracking stopped successfully")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
