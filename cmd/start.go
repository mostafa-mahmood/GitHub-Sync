/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
		pidString := strconv.Itoa(backgroundCmd.Process.Pid)
		err = os.WriteFile("./config/ghs.pid", []byte(pidString), 0644)
		if err != nil {
			fmt.Println("⚠️ Warning: Could not save process ID:", err)
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
			internal.PeriodicCheck()
			// Sleep for 5 minutes before checking again
			time.Sleep(5 * time.Minute)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(timerCmd)
}
