package utils

import (
	"encoding/json"
	"fmt"
	"os"
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

	if state := Exists("./repo"); !state {
		err := os.Mkdir("./repo", 0777)

		if err != nil {
			return fmt.Errorf("error creating repo directory: %v", err)
		}
	}

	if state := Exists("./config"); !state {
		err := os.Mkdir("./config", 0777)

		if err != nil {
			return fmt.Errorf("error creating config directory: %v", err)
		}
	}

	return nil
}

func CreateConfigFiles() error {

	if state := Exists("./config/config.json"); !state {
		file, err := os.Create("./config/config.json")

		if err != nil {
			return fmt.Errorf("error creating config.json: %v", err)
		}

		file.Close()
	}

	if state := Exists("./config/timer.json"); !state {
		file, err := os.Create("./config/timer.json")

		if err != nil {
			return fmt.Errorf("error creating timer.json:: %v", err)
		}

		file.Close()
	}

	return nil
}

// puts json default values in config files
func WriteConfigDefaults() error {
	if IsEmpty("./config/config.json") {
		configDefault := ConfigJson{
			GithubPAT:       "",
			Activity:        "",
			CommitFrequency: 0,
		}

		if err := writeJSON("./config/config.json", configDefault); err != nil {
			return fmt.Errorf("error writing defaults to config.json: %v", err)
		}
	}

	if IsEmpty("./config/timer.json") {
		timerDefault := TimerJson{
			TrackedMinutes:      0,
			TotalSessionMinutes: 0,
			NumberOfCommits:     0,
			LastUpdate:          "",
		}

		if err := writeJSON("./config/timer.json", timerDefault); err != nil {
			return fmt.Errorf("error writing defaults to timer.json: %v", err)
		}
	}

	return nil
}

func WriteTimerDefaults() error {
	if IsEmpty("./config/timer.json") {
		timerDefault := TimerJson{
			TrackedMinutes:      0,
			TotalSessionMinutes: 0,
			NumberOfCommits:     0,
			LastUpdate:          "",
		}

		if err := writeJSON("./config/timer.json", timerDefault); err != nil {
			return fmt.Errorf("error writing defaults to timer.json: %v", err)
		}
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
