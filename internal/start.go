package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mostafa-mahmood/GitHub-Sync/utils"
)

func StartTracking() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🔄 Checking Config Files...")

	if err := utils.CreateConfigDirectories(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	if err := utils.CreateConfigFiles(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	if err := utils.WriteConfigDefaults(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	pat, err := utils.GetPAT()
	pat = strings.TrimSpace(pat)
	if err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}

	if pat == "" {
		fmt.Print("🔑 Enter Your GitHub Personal Access Token: ")
		pat, _ = reader.ReadString('\n')
		pat = pat[:len(pat)-1] // Remove the newline character

		valid, err := IsPatValid(pat)
		if err != nil {
			fmt.Println("❌", err)
			os.Exit(1)
		}
		for !valid {
			if !valid {
				fmt.Print("❌ Invalid PAT. Please enter a valid one: ")
			}

			pat, _ = reader.ReadString('\n')
			pat = pat[:len(pat)-1] // Remove the newline character

			valid, err = IsPatValid(pat)

			if err != nil {
				fmt.Println("❌", err)
				os.Exit(1)
			}
		}

		if err := utils.WritePAT(pat); err != nil {
			fmt.Println("❌", err)
			os.Exit(1)
		}

		fmt.Println("✅ PAT Saved Successfully!")
	} else {
		// Validate the existing PAT
		fmt.Print("🔍 Validating Existing PAT")
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Print(".")
		}
		fmt.Println()

		valid, err := IsPatValid(pat)
		if err != nil {
			fmt.Println("❌", err)
			os.Exit(1)
		}

		if !valid {
			fmt.Println("❌ PAT Is No Longer Valid. Please Enter A Valid One: ")
			pat, _ = reader.ReadString('\n')
			pat = pat[:len(pat)-1] // Remove the newline character
		}

		fmt.Println("✅ PAT Is Still Valid")
	}
	fmt.Println("🔄 Checking Activties Repo...")

	var repoExists bool
	repoExists, err = SpecialRepoExists(pat)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !repoExists {
		err = CreateSpecialRepo(pat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if repoCloned := IsRepoCloned(); !repoCloned {
		username, err := GetGitHubUsername(pat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = CloneRepo(username, pat)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = EnsureRepoFiles()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var editorOpened bool
	var editor string
	editorOpened, editor, err = IsEditorOpened()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !editorOpened {
		fmt.Println("❌ No Text Editors Detected")
		os.Exit(1)
	} else {
		fmt.Printf("Editor Detected: [%v]\n", editor)
	}

	var activity string

	fmt.Print("📝 What are you working on today? ")
	activity, _ = reader.ReadString('\n')
	activity = strings.TrimRight(activity, "\r\n")
	err = utils.WriteActivity(activity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var commitFrequency int
	fmt.Print("🕒 Enter Commit Frequency \n (How Often Do You Want To Push | minimum 100 min): ")
	fmt.Scanln(&commitFrequency)

	for commitFrequency < 100 {
		fmt.Print("❌ Minimum is 100: ")
		fmt.Scanln(&commitFrequency)
	}

	err = utils.WriteCommitFrequency(commitFrequency)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	err = utils.WriteTimerDefaults()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
