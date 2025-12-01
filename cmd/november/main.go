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

	userStats := api.GetUserStats(usrname, rawResponse)

	//debug
	api.Info(userStats)

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
