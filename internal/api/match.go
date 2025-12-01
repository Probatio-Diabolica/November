package api

import (
	"fmt"
	"november/internal/model"
)

func GetRatings(rawResponse model.LeetCodeResponse, userStats model.UserStats) model.UserStats {
	userStats.ContestCount = rawResponse.Data.UserRanking.ContestCount
	userStats.TotalParticipants = rawResponse.Data.UserRanking.TotalParticipants
	userStats.Rating = rawResponse.Data.UserRanking.Rating
	userStats.TopPercent = rawResponse.Data.UserRanking.TopPercent
	userStats.GlobalRanking = rawResponse.Data.UserRanking.GlobalRank
	return userStats
}

func GetUserStats(userName string, rawResponse model.LeetCodeResponse) model.UserStats {
	var userStats model.UserStats
	userStats.UserName = userName

	for _, entry := range rawResponse.Data.MatchedUser.SubmitStats.AcSubmissionNum {
		switch entry.Difficulty {
		case "Easy":
			userStats.EasySolved = entry.Count
		case "Medium":
			userStats.MediumSolved = entry.Count
		case "Hard":
			userStats.HardSolved = entry.Count
		case "All":
			userStats.TotalSolved = entry.Count

		}
	}

	return GetRatings(rawResponse, userStats)
}

// debug
func Info(usr model.UserStats) {
	fmt.Println(" Usr Name->", usr.UserName)
	fmt.Println(" totl ->", usr.TotalSolved)
	fmt.Println(" ez->", usr.EasySolved)
	fmt.Println(" md->", usr.MediumSolved)
	fmt.Println(" hrd->", usr.HardSolved)
	fmt.Println(" rnk->", usr.Ranking)
	fmt.Println(" rating->", usr.Rating)
	fmt.Println(" glb rnk->", usr.GlobalRanking)
	fmt.Println(" CC->", usr.ContestCount)
	fmt.Println(" totl parts ->", usr.TotalParticipants)
	fmt.Println(" tp percent->", usr.TopPercent)

}
