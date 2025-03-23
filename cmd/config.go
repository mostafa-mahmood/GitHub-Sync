/*
Copyright © 2025 mostafa-mahmood
*/
package cmd

import (
	"fmt"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
	"github.com/spf13/cobra"
)

var pat string
var activity string
var frequency int

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update Configuration Data",
	Long: `Use Flags to Change Specific Configuration Properties
like PAT, Activity, or Commit Frequency:
  --pat=<your_personal_access_token>
  --activity=<your_activity>
  --commitFrequency=<minutes>`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flags().NFlag() == 0 {
			currentActivity, err := utils.GetActivity()
			if err != nil {
				fmt.Println(err)
			}

			var currentFrequency int
			currentFrequency, err = utils.GetCommitFrequency()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("📋 Current Configurations \n🔹 Activity: %v \n🔹 Commit Frequency: %v",
				currentActivity, currentFrequency)
		}

		if cmd.Flags().Changed("pat") {
			err := utils.WritePAT(pat)
			if err != nil {
				fmt.Println("❌ Failed to update PAT:", err)
				return
			}
			fmt.Println("✅ PAT updated successfully.")
		}

		if cmd.Flags().Changed("activity") {
			err := utils.WriteActivity(activity)
			if err != nil {
				fmt.Println("❌ Failed to update activity:", err)
				return
			}
			fmt.Println("✅ activity updated successfully.")
		}

		if cmd.Flags().Changed("frequency") {
			err := utils.WriteCommitFrequency(frequency)
			if err != nil {
				fmt.Println("❌ Failed to update frequency:", err)
				return
			}
			fmt.Println("✅ Commit Frequency updated successfully.")
		}

	},
}

func init() {
	configCmd.Flags().StringVar(&pat, "pat", "", "Set your GitHub Personal Access Token")
	configCmd.Flags().StringVar(&activity, "activity", "", "Set your current activity")
	configCmd.Flags().IntVar(&frequency, "frequency", 100, "Set commit frequency in minutes")

	rootCmd.AddCommand(configCmd)
}
