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

// GetBaseDir returns the base directory for all configuration
func GetBaseDir() (string, error) {
	// Get executable path
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error getting executable path: %v", err)
	}

	// Get directory containing the executable
	execDir := filepath.Dir(execPath)
	return execDir, nil
}

// GetConfigDir returns the path to the config directory
func GetConfigDir() (string, error) {
	baseDir, err := GetBaseDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, "config"), nil
}

// GetRepoDir returns the path to the repo directory
func GetRepoDir() (string, error) {
	baseDir, err := GetBaseDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, "repo"), nil
}

func CreateConfigDirectories() error {
	// Get repo directory path
	repoDir, err := GetRepoDir()
	if err != nil {
		return err
	}

	if state := Exists(repoDir); !state {
		err := os.MkdirAll(repoDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating repo directory: %v", err)
		}
	}

	// Get config directory path
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if state := Exists(configDir); !state {
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating config directory: %v", err)
		}
	}

	return nil
}

func CreateConfigFiles() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.json")
	if state := Exists(configPath); !state {
		file, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("error creating config.json: %v", err)
		}
		file.Close()
	}

	timerPath := filepath.Join(configDir, "timer.json")
	if state := Exists(timerPath); !state {
		file, err := os.Create(timerPath)
		if err != nil {
			return fmt.Errorf("error creating timer.json: %v", err)
		}
		file.Close()
	}

	return nil
}

func WriteConfigDefaults() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.json")
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

	timerPath := filepath.Join(configDir, "timer.json")
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
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	timerDefault := TimerJson{
		TrackedMinutes:      0,
		TotalSessionMinutes: 0,
		NumberOfCommits:     0,
		LastUpdate:          "",
	}
	timerPath := filepath.Join(configDir, "timer.json")
	if err := writeJSON(timerPath, timerDefault); err != nil {
		return fmt.Errorf("error writing defaults to timer.json: %v", err)
	}
	return nil
}

func GetConfigFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

func GetTimerFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "timer.json"), nil
}

func GetPidFilePath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "ghs.pid"), nil
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
