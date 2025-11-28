package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const endpoint = "https://leetcode.com/graphql"

const query = `
query getUserProfile($username: String!) {
  matchedUser(username: $username) {
    username
    submitStats {
      acSubmissionNum {
        difficulty
        count
      }
    }
  }
}
`

type Payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type Response struct {
	Data struct {
		MatchedUser struct {
			Username    string `json:"username"`
			SubmitStats struct {
				AcSubmissionNum []struct {
					Difficulty string `json:"difficulty"`
					Count      int    `json:"count"`
				} `json:"acSubmissionNum"`
			} `json:"submitStats"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: november <leetcode_username>")
		return
	}

	username := os.Args[1]

	payload := Payload{
		Query: query,
		Variables: map[string]interface{}{
			"username": username,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var result Response
	json.Unmarshal(raw, &result)

	user := result.Data.MatchedUser

	fmt.Printf("\nUser: %s\n", user.Username)
	fmt.Println("Solved Problems:")

	for _, stat := range user.SubmitStats.AcSubmissionNum {
		fmt.Printf("  %-7s : %d\n", stat.Difficulty, stat.Count)
	}
}
