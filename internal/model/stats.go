package model

type UserStats struct {
	UserName          string
	TotalSolved       int
	EasySolved        int
	MediumSolved      int
	HardSolved        int
	Ranking           int // that useless rank
	GlobalRanking     int // your contest ranking
	ContestCount      int
	TotalParticipants int

	Rating     float32
	TopPercent float32
}

type DifficultyCount struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}

type Ranking struct {
	ContestCount      int     `json:"attendedContestsCount"`
	Rating            float32 `json:"rating"`
	GlobalRank        int     `json:"globalRanking"`
	TotalParticipants int     `json:"totalParticipants"`
	TopPercent        float32 `json:"topPercentage"`
}

type LeetCodeResponse struct {
	Data struct {
		MatchedUser struct {
			Username    string `json:"username"`
			SubmitStats struct {
				AcSubmissionNum []DifficultyCount `json:"acSubmissionNum"`
			} `json:"submitStats"`
		} `json:"matchedUser"`

		UserRanking Ranking `json:"userContestRanking"`
	} `json:"data"`
}
