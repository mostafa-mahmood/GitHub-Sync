package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TimerJson struct {
	TrackedMinutes      int    `json:"tracked_minutes"`
	TotalSessionMinutes int    `json:"total_session_minutes"`
	NumberOfCommits     int    `json:"number_of_commits"`
	LastUpdate          string `json:"last_update"`
}

type ConfigJson struct {
	GithubPAT       string `json:"Github_Personal_Access_Token"`
	Activity        string `json:"Activity"`
	CommitFrequency int    `json:"Commit_Frequency"`
}

func CreateConfigDirectories() error {
	if state := Exists(filepath.Join(".", "repo")); !state {
		err := os.Mkdir(filepath.Join(".", "repo"), 0755)
		if err != nil {
			return fmt.Errorf("error creating repo directory: %v", err)
		}
	}

	if state := Exists(filepath.Join(".", "config")); !state {
		err := os.Mkdir(filepath.Join(".", "config"), 0755)
		if err != nil {
			return fmt.Errorf("error creating config directory: %v", err)
		}
	}

	return nil
}

func CreateConfigFiles() error {
	configPath := filepath.Join(".", "config", "config.json")
	if state := Exists(configPath); !state {
		file, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("error creating config.json: %v", err)
		}
		file.Close()
	}

	timerPath := filepath.Join(".", "config", "timer.json")
	if state := Exists(timerPath); !state {
		file, err := os.Create(timerPath)
		if err != nil {
			return fmt.Errorf("error creating timer.json: %v", err)
		}
		file.Close()
	}

	return nil
}

// puts json default values in config files
func WriteConfigDefaults() error {
	configPath := filepath.Join(".", "config", "config.json")
	if IsEmpty(configPath) {
		configDefault := ConfigJson{
			GithubPAT:       "",
			Activity:        "",
			CommitFrequency: 0,
		}

		if err := writeJSON(configPath, configDefault); err != nil {
			return fmt.Errorf("error writing defaults to config.json: %v", err)
		}
	}

	timerPath := filepath.Join(".", "config", "timer.json")
	if IsEmpty(timerPath) {
		timerDefault := TimerJson{
			TrackedMinutes:      0,
			TotalSessionMinutes: 0,
			NumberOfCommits:     0,
			LastUpdate:          "",
		}

		if err := writeJSON(timerPath, timerDefault); err != nil {
			return fmt.Errorf("error writing defaults to timer.json: %v", err)
		}
	}

	return nil
}

func WriteTimerDefaults() error {
	timerDefault := TimerJson{
		TrackedMinutes:      0,
		TotalSessionMinutes: 0,
		NumberOfCommits:     0,
		LastUpdate:          "",
	}
	timerPath := filepath.Join(".", "config", "timer.json")
	if err := writeJSON(timerPath, timerDefault); err != nil {
		return fmt.Errorf("error writing defaults to timer.json: %v", err)
	}
	return nil
}

func writeJSON(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func IsEmpty(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return info.Size() == 0
}

func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
