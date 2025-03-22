package internal

import (
	"fmt"
	"os"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

func PeriodicCheck() {

	editorOpened, _, err := IsEditorOpened()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	if editorOpened {
		err := utils.UpdateTrackedMinutes()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		err = utils.UpdateTotalSessionMinutes()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		err = utils.UpdateLastUpdate()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		var trackedMinutes int
		trackedMinutes, err = utils.GetTrackedMinutes()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		var commitFrequency int
		commitFrequency, err = utils.GetCommitFrequency()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		if trackedMinutes >= commitFrequency {
			err = utils.ResetTrackedMinutes()

			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			commitMessage, err := FormatCommitMessage()
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			err = AppendToLog(commitMessage)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			err = CommitAndPushChanges(commitMessage)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
	}
}
