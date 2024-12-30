package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	username, err := getUsernameFromCli()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	response, err := http.Get(buildGitHubUrl(username))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	// Parse the response
	githubEvents := []GitHubEvent{}
	json.Unmarshal(body, &githubEvents)

	printGitHubActivity(githubEvents)
}
