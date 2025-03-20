package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Handels git operations: clone, commit, push

func IsRepoCloned() bool {
	_, err := os.Stat("repo/.git")
	return !os.IsNotExist(err)
}

func CloneRepo(username, PAT string) error {
	if IsRepoCloned() {
		fmt.Println("Repository already cloned.")
		return nil
	}

	url := fmt.Sprintf("https://%s@github.com/%s/Activities.git", PAT, username)

	cmd := exec.Command("git", "clone", url, "repo")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone repo: %v", err)
	}

	fmt.Println("Repository cloned successfully!")
	return nil
}

func CommitAndPushChanges(message string) error {
	cmds := [][]string{
		{"git", "-C", "repo", "add", "."},
		{"git", "-C", "repo", "commit", "-m", message},
		{"git", "-C", "repo", "push"},
	}

	for _, cmdArgs := range cmds {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error running %s: %v", cmdArgs[0], err)
		}
	}

	fmt.Println("Changes committed and pushed successfully!")
	return nil
}
