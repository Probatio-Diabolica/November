package main

import (
	"fmt"
	"log"
	"november/internal/api"
	"os"
)

type userStats struct{}

func getUsername() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("username not provided")
	}

	return os.Args[1], nil
}

func showData(usrData userStats) {
	fmt.Println("displayed")
}

func fetchData(usrname string) (userStats, error) {

	stats, err := api.FetchUserStats(usrname)
	if err != nil {
		return userStats{}, fmt.Errorf("error getting this")
	}
	for _, val := range stats {
		fmt.Printf("%c", val)
	}

	return userStats{}, nil
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
