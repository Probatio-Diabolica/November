package main

import (
	"encoding/json"
	"fmt"
	"log"
	"november/internal/api"
	"november/internal/model"
	"os"
)

func getUsername() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("username not provided")
	}

	return os.Args[1], nil
}

func showData(usrData model.UserStats) {
	return
}

func fetchData(usrname string) (model.UserStats, error) {

	fmt.Println(usrname)
	rawStats, err := api.FetchUserStats(usrname)
	if err != nil {
		return model.UserStats{}, fmt.Errorf("error getting this")
	}

	// js for debugging
	// for _, val := range rawStats {
	// 	fmt.Printf("%c", val)
	// }

	var rawResponse model.LeetCodeResponse
	err = json.Unmarshal(rawStats, &rawResponse)
	if err != nil {
		return model.UserStats{}, err
	}

	json.Unmarshal(rawStats, &rawResponse)

	var userStats model.UserStats
	userStats.UserName = usrname

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

	fmt.Println(userStats.TotalSolved)
	return userStats, nil
}

func main() {

	username, err := getUsername()
	if err != nil {
		log.Fatalln("usage : nov <username>")
	}

	usrData, err := fetchData(username)
	if err != nil {
		log.Fatalln("wrong username")
	}
	showData(usrData)

}
