package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

func FormatCommitMessage() (string, error) {
	var err error

	var workSession int
	workSession, err = utils.GetTotalSessionMinutes()
	if err != nil {
		return "", err
	}
	workSession /= 60

	var activity string
	activity, err = utils.GetActivity()
	if err != nil {
		return "", err
	}

	_, editor, err := IsEditorOpened()
	if err != nil {
		return "", err
	}

	now := time.Now()
	date := now.Format("2006-01-02")
	timeStr := now.Format("15:04:05")

	commitMessage := fmt.Sprintf("Work Session: %vhr | Activity: %v | Editor Used: %v | Date: %v | Time: %v",
		workSession, activity, editor, date, timeStr)

	return commitMessage, nil
}

func AppendToLog(message string) error {
	filePath := filepath.Join(".", "repo", "Activities", "log.txt")

	if repoCloned, _ := IsRepoCloned(); !repoCloned {
		return fmt.Errorf("can't append repository is not cloned")
	}

	if fileExists := utils.Exists(filePath); !fileExists {
		return fmt.Errorf("can't append log.txt doesn't exists")
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + message + "\n")
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
