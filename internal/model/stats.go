package model

type UserStats struct {
	UserName     string
	TotalSolved  int
	EasySolved   int
	MediumSolved int
	HardSolved   int
	Ranking      int
	Rating       int
}

type DifficultyCount struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}
type LeetCodeResponse struct {
	Data struct {
		MatchedUser struct {
			Username    string `json:"username"`
			SubmitStats struct {
				AcSubmissionNum []DifficultyCount `json:"acSubmissionNum"`
			} `json:"submitStats"`
		} `json:"matchedUser"`
	} `json:"data"`
}
