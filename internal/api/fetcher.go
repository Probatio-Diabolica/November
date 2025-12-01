package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const userStatsQuery = `
query getUserStats($username: String!) {
	matchedUser(username: $username) {
		submitStats {
			acSubmissionNum {
				difficulty
				count
			}
		}
	}
	userContestRanking(username: $username){
		attendedContestsCount
		rating
		globalRanking
		totalParticipants
		topPercentage
	}
}
`

const url = "https://leetcode.com/graphql"

type PostQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func GetData(query string, variables map[string]interface{}) ([]byte, error) {

	body := PostQuery{
		Query:     query,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("ight json marshal failed </3")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("sadge! post req failed")
	}

	req.Header.Set("Content-Type", "application/json")

	//lmaoo i forgor to setup client
	client := &http.Client{}

	//yea then do this client
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("req failed")
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func FetchUserStats(username string) ([]byte, error) {
	vars := map[string]interface{}{
		"username": username,
	}

	return GetData(userStatsQuery, vars)
}
