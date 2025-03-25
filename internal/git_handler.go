package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

// Handles git operations: clone, commit, push

func IsRepoCloned() (bool, error) {
	repoDir, err := utils.GetRepoDir()
	if err != nil {
		return false, fmt.Errorf("error getting repo directory: %v", err)
	}

	_, err = os.Stat(filepath.Join(repoDir, "Activities", ".git"))
	return !os.IsNotExist(err), nil
}

func CloneRepo(username, PAT string) error {
	PAT = strings.TrimSpace(PAT)

	repoDir, err := utils.GetRepoDir()
	if err != nil {
		return fmt.Errorf("error getting repo directory: %v", err)
	}

	repoActivitiesPath := filepath.Join(repoDir, "Activities")

	// Check if "repo" exists but is a file
	if utils.Exists(repoDir) && !utils.IsDirectory(repoDir) {
		err := os.Remove(repoDir) // Delete the file
		if err != nil {
			return fmt.Errorf("failed to remove existing file 'repo': %v", err)
		}
	}

	// Ensure "repo" directory exists
	if !utils.Exists(repoActivitiesPath) {
		err := os.MkdirAll(repoDir, 0755) // Use MkdirAll to create "repo" if missing
		if err != nil {
			return fmt.Errorf("failed to create repo directory: %v", err)
		}
	}

	isCloned, err := IsRepoCloned()
	if err != nil {
		return err
	}

	if isCloned {
		return nil
	}

	url := fmt.Sprintf("https://%s@github.com/%s/Activities.git", PAT, username)

	cmd := exec.Command("git", "clone", url, repoActivitiesPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone repo: %v", err)
	}

	fmt.Println("Repository cloned successfully inside 'repo/Activities/'!")
	return nil
}

func CommitAndPushChanges(message string) error {
	repoDir, err := utils.GetRepoDir()
	if err != nil {
		return fmt.Errorf("error getting repo directory: %v", err)
	}

	repoActivitiesPath := filepath.Join(repoDir, "Activities")

	cmds := [][]string{
		{"git", "-C", repoActivitiesPath, "add", "--all"},
		{"git", "-C", repoActivitiesPath, "commit", "-m", message},
		{"git", "-C", repoActivitiesPath, "push", "-f"},
	}

	for _, cmdArgs := range cmds {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error running %s: %v", strings.Join(cmdArgs, " "), err)
		}
	}

	fmt.Println("Changes committed and pushed successfully!")
	return nil
}

func EnsureRepoFiles() error {
	repoDir, err := utils.GetRepoDir()
	if err != nil {
		return fmt.Errorf("error getting repo directory: %v", err)
	}

	repoActivitiesPath := filepath.Join(repoDir, "Activities")

	// Ensure the repo directory exists
	if !utils.Exists(repoActivitiesPath) {
		err := os.MkdirAll(repoActivitiesPath, 0755) // Create the directory if missing
		if err != nil {
			return fmt.Errorf("failed to create repo directory: %v", err)
		}
	}

	// Check and create README.md
	readmePath := filepath.Join(repoActivitiesPath, "README.md")
	if !utils.Exists(readmePath) {
		content := []byte(`
### Hey there, fellow programmer! 👋  

## What is this?
This repository is automatically managed by [GitHub-Sync](https://github.com/mostafa-mahmood/GitHub-Sync) (ghs).      

This repository acts as a logbook for my coding sessions. It keeps track of when and how long I’ve been coding,   
along with details like the editor used and a description of current activity.   
Logging your session details into log.txt at regular intervals.
`)

		err := os.WriteFile(readmePath, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to create README.md: %v", err)
		}
	}

	logPath := filepath.Join(repoActivitiesPath, "log.txt")
	if !utils.Exists(logPath) {
		err := os.WriteFile(logPath, []byte(""), 0644)
		if err != nil {
			return fmt.Errorf("failed to create log.json: %v", err)
		}
	}

	return nil
}
