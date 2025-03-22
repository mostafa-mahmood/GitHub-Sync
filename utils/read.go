package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadTimer() (TimerJson, error) {
	var timer TimerJson

	data, err := os.ReadFile("./config/timer.json")
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

	data, err := os.ReadFile("./config/config.json")
	if err != nil {
		return config, fmt.Errorf("error reading config.json: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config.json: %v", err)
	}

	return config, nil
}
