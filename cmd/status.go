/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"github.com/mostafa-mahmood/GitHub-Sync/internal"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Provides insights about current tracking status.",
	Long:  `Provides insights about current tracking status.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
