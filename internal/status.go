package internal

import (
	"fmt"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

func PrintStatus() error {
	var trackedWork int
	trackedWork, err := utils.GetTotalSessionMinutes()
	if err != nil {
		return err
	}

	var lastUpdate string
	lastUpdate, err = utils.GetLastUpdate()
	if err != nil {
		return err
	}

	var trackedMinutes int
	trackedMinutes, err = utils.GetTrackedMinutes()
	if err != nil {
		return err
	}

	var CommitFrequency int
	CommitFrequency, err = utils.GetCommitFrequency()
	if err != nil {
		return err
	}
	minutesRemainning := CommitFrequency - trackedMinutes

	var editorOpened bool
	var editor string
	editorOpened, editor, err = IsEditorOpened()
	if err != nil {
		return err
	}

	var statusMessage string
	if editorOpened {
		statusMessage = fmt.Sprintf("📈 Tracker Status \n🔵 Tracked Work: %v \n🔵 Last Update: %v \n🔵 Time Remaining Before Next Update: %v \n🔵 Editor: %v(Running)",
			trackedWork, lastUpdate, minutesRemainning, editor)
	} else {
		statusMessage = fmt.Sprintf("📈Tracker Status \n🔵 Tracked Work: %v \n🔵 Last Update: %v \n🔵 Time Remaining Before Next Update: %v \n🔵 Editor: No Editor Detected",
			trackedWork, lastUpdate, minutesRemainning)
	}

	fmt.Println(statusMessage)
	return nil
}
