/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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

func init() {
	rootCmd.AddCommand(stopCmd)
}
