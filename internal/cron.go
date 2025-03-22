package internal

import (
	"fmt"
	"os"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

func PeriodicCheck() {
	editorOpened, _, err := IsEditorOpened()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking editor status: %v\n", err)
		os.Exit(1)
	}

	if editorOpened {
		err := utils.UpdateTrackedMinutes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating tracked minutes: %v\n", err)
			os.Exit(1)
		}

		err = utils.UpdateTotalSessionMinutes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating total session minutes: %v\n", err)
			os.Exit(1)
		}

		err = utils.UpdateLastUpdate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating last update: %v\n", err)
			os.Exit(1)
		}

		var trackedMinutes int
		trackedMinutes, err = utils.GetTrackedMinutes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting tracked minutes: %v\n", err)
			os.Exit(1)
		}

		var commitFrequency int
		commitFrequency, err = utils.GetCommitFrequency()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting commit frequency: %v\n", err)
			os.Exit(1)
		}

		if trackedMinutes >= commitFrequency {
			err = utils.ResetTrackedMinutes()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error resetting tracked minutes: %v\n", err)
				os.Exit(1)
			}

			commitMessage, err := FormatCommitMessage()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting commit message: %v\n", err)
				os.Exit(1)
			}

			err = AppendToLog(commitMessage)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error appending to log: %v\n", err)
				os.Exit(1)
			}

			err = CommitAndPushChanges(commitMessage)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error committing and pushing changes: %v\n", err)
				os.Exit(1)
			}

			err = utils.UpdateNumberOfCommits()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error updating number of commits: %v\n", err)
				os.Exit(1)
			}
		}
	}
}
