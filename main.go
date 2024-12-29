package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type github_response struct {
	id string
}

func main() {
	// Get the CLI arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No username provided. Please provide a username")
		os.Exit(0)
	}

	username := args[0]

	fmt.Println("Username: ", username)

	// Fetch Data from GitHub – https://api.github.com/users/<username>/events
	response, err := http.Get(builtGitHubUrl(username))

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
	var result []map[string]interface{}
	json.Unmarshal(body, &result)

	fmt.Println(result)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
		fmt.Println()
	}

	// Print the output. Example:
	//  Output:
	// - Pushed 3 commits to kamranahmedse/developer-roadmap
	// - Opened a new issue in kamranahmedse/developer-roadmap
	// - Starred kamranahmedse/developer-roadmap
	// - ...
}

func builtGitHubUrl(username string) string {
	return "https://api.github.com/users/" + username + "/events"
}
