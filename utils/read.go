package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadTimer() (TimerJson, error) {
	var timer TimerJson

	timerPath, err := GetTimerFilePath()
	if err != nil {
		return timer, fmt.Errorf("error getting timer.json path: %v", err)
	}

	data, err := os.ReadFile(timerPath)
	if err != nil {
		return timer, fmt.Errorf("error reading timer.json: %v", err)
	}

	err = json.Unmarshal(data, &timer)
	if err != nil {
		return timer, fmt.Errorf("error parsing timer.json: %v", err)
	}

	return timer, nil
}

func ReadConfig() (ConfigJson, error) {
	var config ConfigJson

	configPath, err := GetConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("error getting config.json path: %v", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("error reading config.json: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config.json: %v", err)
	}

	return config, nil
}
