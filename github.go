package main

import (
	"fmt"
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
	Action   string         `json:"action"`
	Comments PayloadComment `json:"comment"`
	Issue    PayloadIssue   `json:"issue"`
	Pages    []PayloadPage  `json:"pages"`
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

type Repository struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func buildGitHubUrl(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events", username)
}
