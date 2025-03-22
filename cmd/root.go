/*
Copyright © 2025 mostafa-mahmood
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghs",
	Short: "GitHub-Sync: A Tool That Tracks Coding Activity And Syncs It With GitHub",
	Long: `GitHub-Sync (ghs) A Tool That Tracks Your Coding Activity and Syncs It With GitHub.
	
- Detects when your editor is open
- Tracks coding time and commits at set intervals
- Stores logs of commit history`,
}

// Execute runs the CLI and handles errors.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Placeholder: We will load config in the future
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.ghs.yaml)")
}
