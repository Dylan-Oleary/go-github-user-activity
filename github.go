package main

import (
	"fmt"
	"strings"
	"time"
)

type GitHubEvent struct {
	ID         string     `json:"id"`
	EventType  string     `json:"type"`
	Actor      Actor      `json:"actor"`
	Repository Repository `json:"repo"`
	Payload    Payload    `json:"payload"`
	Public     bool       `json:"public"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Actor struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Payload struct {
	Action      string             `json:"action"`
	Comments    PayloadComment     `json:"comment"`
	Commits     []PayloadCommit    `json:"commits"`
	Issue       PayloadIssue       `json:"issue"`
	Pages       []PayloadPage      `json:"pages"`
	PullRequest PayloadPullRequest `json:"pull_request"`
	Ref         string             `json:"ref"`
	RefType     string             `json:"ref_type"`
}

type PayloadComment struct {
	ID                string    `json:"id"`
	AuthorAssociation string    `json:"author_association"`
	Body              string    `json:"body"`
	BodyHtml          string    `json:"body_html"`
	BodyText          string    `json:"body_text"`
	HtmlUrl           string    `json:"html_url"`
	IssueUrl          string    `json:"issue_url"`
	NodeID            string    `json:"node_id"`
	Url               string    `json:"url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type PayloadCommit struct {
	ID  string `json:"id"`
	Sha string `json:"sha"`
}

type PayloadIssue struct {
	ID            string `json:"id"`
	Body          string `json:"body"`
	CommentsUrl   string `json:"comments_url"`
	EventsUrl     string `json:"events_url"`
	HtmlUrl       string `json:"html_url"`
	RepositoryUrl string `json:"repository_url"`
	LabelsUrl     string `json:"labels_url"`
	NodeID        string `json:"node_id"`
	Number        int    `json:"number"`
	State         string `json:"state"`
	StateReason   string `json:"state_reason"`
	Title         string `json:"title"`
	Url           string `json:"url"`
}

type PayloadPage struct {
	Action   string `json:"action"`
	HtmlUrl  string `json:"html_url"`
	PageName string `json:"page_name"`
	Sha      string `json:"sha"`
	Summary  string `json:"summary"`
	Title    string `json:"title"`
}

type PayloadPullRequest struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
	Url    string `json:"url"`
}

type Repository struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func buildGitHubUrl(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", username)
}

func printGitHubActivity(events []GitHubEvent) {
	fmt.Printf("Output:\n\n")

	for i := 0; i < len(events); i++ {
		event := events[i]

		ref := event.Payload.Ref
		refType := event.Payload.RefType
		repository := event.Repository

		// This is not inclusive of all GitHub event types
		switch event.EventType {
		case "CreateEvent":
			entity := ""

			if refType == "repository" {
				entity = refType
			} else {
				entity = fmt.Sprintf("%s (%s)", refType, ref)
			}

			fmt.Printf("- Created new %s in %s\n\n", entity, repository.Name)
			break
		case "DeleteEvent":
			entity := ""

			if refType == "repository" {
				entity = refType
			} else {
				entity = fmt.Sprintf("%s (%s)", refType, ref)
			}

			fmt.Printf("- Deleted %s in %s\n\n", entity, repository.Name)
			break
		case "PullRequestEvent":
			action := fmt.Sprintf("%s%s", strings.ToUpper(event.Payload.Action[0:1]), event.Payload.Action[1:])

			fmt.Printf("- %s pull request #%d in %s\n\n", action, event.Payload.PullRequest.Number, repository.Name)
			break
		case "PushEvent":
			fmt.Printf("- Pushed %d commit(s) to %s\n\n", len(event.Payload.Commits), repository.Name)
			break
		}
	}
}
