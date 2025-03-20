package utils

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

func CreateConfigDirectories() {

}

func CreateConfigFiles() {

}

func Exists(path string) {

}
