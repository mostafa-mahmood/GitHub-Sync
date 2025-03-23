/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the tracked time and commit history",
	Long: `Resets the timer.json file, clearing tracked minutes, session time, commit count, and last update.
Your GitHub API token and settings will remain unchanged.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Resetting tracked time...")

		if err := utils.WriteTimerDefaults(); err != nil {
			fmt.Println("❌ Error resetting timer:", err)
		} else {
			fmt.Println("✅ Timer reset successfully.")
		}

		fmt.Println("Reset complete. Tracking will restart fresh.")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
