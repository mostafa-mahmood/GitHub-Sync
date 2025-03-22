package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// checks if the special repo already exists
func SpecialRepoExists(PAT string) (bool, error) {
	PAT = strings.TrimSpace(PAT)
	username, err := GetGitHubUsername(PAT)
	if err != nil {
		return false, err
	}

	URL := fmt.Sprintf("https://api.github.com/repos/%s/Activities", username)

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "token "+PAT)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("GitHub API request failed with status: %d", res.StatusCode)
	}

	return true, nil
}

// creates the special repo the tool will be pushing to
func CreateSpecialRepo(PAT string) error {
	PAT = strings.TrimSpace(PAT)
	data := `{
		"name": "Activities",
		"private": false,
		"auto_init": false
	}`

	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "token "+PAT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		log.Fatalf("Failed: %d - %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	return nil
}

// checks if the provided GitHub Personal Access Token is valid
func IsPatValid(PAT string) (bool, error) {
	PAT = strings.TrimSpace(PAT)
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return false, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+PAT)

	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

// reterives github user name using Personal Access Token
func GetGitHubUsername(PAT string) (string, error) {
	PAT = strings.TrimSpace(PAT)
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "token "+PAT)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch user info, status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	username, ok := result["login"].(string)
	if !ok {
		return "", fmt.Errorf("could not extract username")
	}

	return username, nil
}
