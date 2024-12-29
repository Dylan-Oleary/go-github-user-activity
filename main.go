package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	// Print the output. Example:
	fmt.Println("Output:")

	for i := 0; i < len(githubEvents); i++ {
		event := githubEvents[i]

		switch event.EventType {
		case "CommitCommentEvent":
			fmt.Println("CommitCommentEvent")
			break
		case "CreateEvent":
			fmt.Println("CreateEvent")
			break
		case "DeleteEvent":
			fmt.Println("DeleteEvent")
			break
		case "ForkEvent":
			fmt.Println("ForkEvent")
			break
		case "GollumEvent":
			fmt.Println("GollumEvent")
			break
		case "IssueCommentEvent":
			fmt.Println("IssueCommentEvent")
			break
		case "IssuesEvent":
			fmt.Println("IssuesEvent")
			break
		case "MemberEvent":
			fmt.Println("MemberEvent")
			break
		case "PublicEvent":
			fmt.Println("PublicEvent")
			break
		case "PullRequestEvent":
			fmt.Println("PullRequestEvent")
			break
		case "PullRequestReviewEvent":
			fmt.Println("PullRequestReviewEvent")
			break
		case "PullRequestReviewCommentEvent":
			fmt.Println("PullRequestReviewCommentEvent")
			break
		case "PullRequestReviewThreadEvent":
			fmt.Println("PullRequestReviewThreadEvent")
			break
		case "PushEvent":
			fmt.Println("PushEvent")
			break
		case "ReleaseEvent":
			fmt.Println("ReleaseEvent")
			break
		case "SponsorshipEvent":
			fmt.Println("SponsorshipEvent")
			break
		case "WatchEvent":
			fmt.Println("WatchEvent")
			break
		}
	}
	//  Output:
	// - Pushed 3 commits to dylan-oleary/developer-roadmap
	// - Opened a new issue in dylan-oleary/developer-roadmap
	// - Starred dylan-oleary/developer-roadmap
	// - ...
}
