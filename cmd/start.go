/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mostafa-mahmood/GitHub-Sync/internal"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts tracking coding activity and pushes every interval",
	Long: `Starts tracking the user's coding activity by detecting open editors.
Once the commit interval is reached, changes are committed and pushed to GitHub.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartTracking()

		// Get the full path to the current executable
		executablePath, err := os.Executable()
		if err != nil {
			fmt.Println("❌ Error getting executable path:", err)
			os.Exit(1)
		}

		// Verify the executable exists
		if _, err := os.Stat(executablePath); os.IsNotExist(err) {
			fmt.Println("❌ Executable not found at:", executablePath)
			os.Exit(1)
		}

		// Launch the timer process in the background
		backgroundCmd := exec.Command(executablePath, "timer")
		err = backgroundCmd.Start()
		if err != nil {
			fmt.Println("❌ Error starting background process:", err)
			os.Exit(1)
		}

		// Save the PID to a file in the config directory
		execDir := filepath.Dir(executablePath)

		// Create a config directory relative to the executable location
		configDir := filepath.Join(execDir, "config")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			fmt.Println("❌ Error creating config directory:", err)
			os.Exit(1)
		}

		// Save the PID to a file in the config directory
		pidPath := filepath.Join(configDir, "ghs.pid")
		pidString := strconv.Itoa(backgroundCmd.Process.Pid)
		err = os.WriteFile(pidPath, []byte(pidString), 0644)
		if err != nil {
			fmt.Println("❌ Error saving process ID:", err)
			os.Exit(1)
		}

		fmt.Println("⏳ Tracking Started, Happy Coding")
	},
}

var timerCmd = &cobra.Command{
	Use:    "timer",
	Short:  "Background process for periodic checks",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			// Sleep for 5 minutes before checking
			time.Sleep(5 * time.Minute)
			internal.PeriodicCheck()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(timerCmd)
}
