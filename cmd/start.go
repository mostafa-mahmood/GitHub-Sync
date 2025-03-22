/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"github.com/mostafa-mahmood/GitHub-Sync/internal"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts tracking coding activity and pushes every interval",
	Long: `Starts tracking the user's coding activity by detecting open editors.
Once the commit interval is reached, changes are committed and pushed to GitHub.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.StartTracking()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
