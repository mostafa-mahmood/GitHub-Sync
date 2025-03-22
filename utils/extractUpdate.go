package utils

import (
	"fmt"
	"time"
)

const (
	configPath = "./config/config.json"
	timerPath  = "./config/timer.json"
)

// UpdateConfig provides a generic way to update the config file
func updateConfig(modifier func(*ConfigJson) error) error {
	configStruct, err := ReadConfig()
	if err != nil {
		return err
	}

	if err := modifier(&configStruct); err != nil {
		return err
	}

	return writeJSON(configPath, configStruct)
}

// UpdateTimer provides a generic way to update the timer file
func updateTimer(modifier func(*TimerJson) error) error {
	timerStruct, err := ReadTimer()
	if err != nil {
		return err
	}

	if err := modifier(&timerStruct); err != nil {
		return err
	}

	return writeJSON(timerPath, timerStruct)
}

// Config operations
func GetPAT() (string, error) {
	configStruct, err := ReadConfig()
	if err != nil {
		return "", err
	}
	return configStruct.GithubPAT, nil
}

func WritePAT(pat string) error {
	return updateConfig(func(c *ConfigJson) error {
		c.GithubPAT = pat
		return nil
	})
}

func GetActivity() (string, error) {
	configStruct, err := ReadConfig()
	if err != nil {
		return "", err
	}
	return configStruct.Activity, nil
}

func WriteActivity(activity string) error {
	return updateConfig(func(c *ConfigJson) error {
		c.Activity = activity
		return nil
	})
}

func GetCommitFrequency() (int, error) {
	configStruct, err := ReadConfig()
	if err != nil {
		return 0, err
	}
	return configStruct.CommitFrequency, nil
}

func WriteCommitFrequency(frequency int) error {
	if frequency < 100 {
		return fmt.Errorf("commit frequency should be at least 100")
	}

	return updateConfig(func(c *ConfigJson) error {
		c.CommitFrequency = frequency
		return nil
	})
}

// Timer operations
func GetTrackedMinutes() (int, error) {
	timerStruct, err := ReadTimer()
	if err != nil {
		return 0, err
	}
	return timerStruct.TrackedMinutes, nil
}

func UpdateTrackedMinutes() error {
	return updateTimer(func(t *TimerJson) error {
		t.TrackedMinutes += 5
		return nil
	})
}

func ResetTrackedMinutes() error {
	return updateTimer(func(t *TimerJson) error {
		t.TrackedMinutes = 0
		return nil
	})
}

func GetTotalSessionMinutes() (int, error) {
	timerStruct, err := ReadTimer()
	if err != nil {
		return 0, err
	}
	return timerStruct.TotalSessionMinutes, nil
}

func UpdateTotalSessionMinutes() error {
	return updateTimer(func(t *TimerJson) error {
		t.TotalSessionMinutes += 5
		return nil
	})
}

func ResetTotalSessionMinutes() error {
	return updateTimer(func(t *TimerJson) error {
		t.TotalSessionMinutes = 0
		return nil
	})
}

func GetNumberOfCommits() (int, error) {
	timerStruct, err := ReadTimer()
	if err != nil {
		return 0, err
	}
	return timerStruct.NumberOfCommits, nil
}

func UpdateNumberOfCommits() error {
	return updateTimer(func(t *TimerJson) error {
		t.NumberOfCommits += 1
		return nil
	})
}

func ResetNumberOfCommits() error {
	return updateTimer(func(t *TimerJson) error {
		t.NumberOfCommits = 0
		return nil
	})
}

func GetLastUpdate() (string, error) {
	timerStruct, err := ReadTimer()
	if err != nil {
		return "", err
	}
	return timerStruct.LastUpdate, nil
}

func UpdateLastUpdate() error {
	return updateTimer(func(t *TimerJson) error {
		t.LastUpdate = fmt.Sprintf("%v", time.Now())
		return nil
	})
}
