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

func IsRepoCloned() bool {
	_, err := os.Stat(filepath.Join("repo", "Activities", ".git"))
	return !os.IsNotExist(err)
}

func CloneRepo(username, PAT string) error {
	PAT = strings.TrimSpace(PAT)
	repoPath := "repo"
	repoActivitiesPath := filepath.Join("repo", "Activities")

	// Check if "repo" exists but is a file
	if utils.Exists(repoPath) && !utils.IsDirectory(repoPath) {
		err := os.Remove(repoPath) // Delete the file
		if err != nil {
			return fmt.Errorf("failed to remove existing file 'repo': %v", err)
		}
	}

	// Ensure "repo" directory exists
	if !utils.Exists(repoActivitiesPath) {
		err := os.MkdirAll(repoPath, 0755) // Use MkdirAll to create "repo" if missing
		if err != nil {
			return fmt.Errorf("failed to create repo directory: %v", err)
		}
	}

	if IsRepoCloned() {
		return nil
	}

	url := fmt.Sprintf("https://%s@github.com/%s/Activities.git", PAT, username)

	cmd := exec.Command("git", "clone", url, repoActivitiesPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone repo: %v", err)
	}

	fmt.Println("Repository cloned successfully inside 'repo/Activities/'! 🎉")
	return nil
}

func CommitAndPushChanges(message string) error {
	repoPath := filepath.Join("repo", "Activities")

	cmds := [][]string{
		{"git", "-C", repoPath, "add", "--all"},
		{"git", "-C", repoPath, "commit", "-m", message},
		{"git", "-C", repoPath, "push", "-f"},
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
	repoPath := filepath.Join("repo", "Activities")

	// Ensure the repo directory exists
	if !utils.Exists(repoPath) {
		err := os.MkdirAll(repoPath, 0755) // Create the directory if missing
		if err != nil {
			return fmt.Errorf("failed to create repo directory: %v", err)
		}
	}

	// Check and create README.md
	readmePath := filepath.Join(repoPath, "README.md")
	if !utils.Exists(readmePath) {
		content := []byte(`
### Hey there, fellow coder! 👋  

## What is this?
Ever coded daily for hours and finally pushed your changes after a while, and GitHub was like: **"Oh, so you only worked today, huh?"**
And it counts as a single contribution for the whole week.  
Yeah, same. GitHub activity tracking is a bit... let's say, "unreliable" (*cough* unfair *cough*)
And the contribution graph might be the only way for others to know your coding activity.

That's where [GitHub-Sync](https://github.com/mostafa-mahmood/GitHub-Sync) comes to the save

## How does this work?
This bad boy keeps track of your local coding sessions. If your editor is open, it counts the minutes.
Once you hit 100 minutes, BAM 💥 — an automatic push to GitHub happens.

## Why?
- So GitHub shows you actually code and don't just show up once a week.  
- So your contribution graph looks like a masterpiece, not a graveyard.

Enjoy, and may your GitHub graph forever shine green! 🌱  
`)

		err := os.WriteFile(readmePath, content, 0644)
		if err != nil {
			return fmt.Errorf("failed to create README.md: %v", err)
		}
	}

	// Check and create log.json
	logPath := filepath.Join(repoPath, "log.txt")
	if !utils.Exists(logPath) {
		err := os.WriteFile(logPath, []byte(""), 0644)
		if err != nil {
			return fmt.Errorf("failed to create log.json: %v", err)
		}
	}

	return nil
}
