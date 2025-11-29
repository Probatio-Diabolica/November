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

type LeetcodeResponse struct {
	Data struct {
		MatchedUser struct {
			SubmitStats struct {
				AcSubmissionNum []struct {
					Difficulty string `json:"difficulty"`
					Count      int    `json:"count"`
				} `json:"acSubmissionNum"`
			} `json:"submitStats"`
		} `json:"matchedUser"`
	} `json:"data"`
}
