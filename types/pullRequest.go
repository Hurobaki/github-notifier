package types

type PullRequestStatus string

type PullRequestReviewStatus string

const (
	Closed        PullRequestStatus = "closed"
	Opened        PullRequestStatus = "opened"
	ReviewRequest PullRequestStatus = "review_requested"
	Labeled       PullRequestStatus = "labeled"
)

const (
	Submitted PullRequestReviewStatus = "submitted"
	Edited    PullRequestReviewStatus = "edited"
	Dismissed PullRequestReviewStatus = "dismissed"
)

const (
	PullRequestEvent       string = "pull_request"
	PullRequestReviewEvent string = "pull_request_review"
)

type PullRequestReview struct {
	Action      PullRequestReviewStatus      `json:"action"`
	Review      PullRequestReviewInformation `json:"review"`
	Information PullRequestInformation       `json:"pull_request"`
	Sender      GithubUser                   `json:"sender"`
}

type PullRequest struct {
	Action      PullRequestStatus      `json:"action"`
	Number      int                    `json:"number"`
	Information PullRequestInformation `json:"pull_request"`
	Sender      GithubUser             `json:"sender"`
}

type PullRequestReviewInformation struct {
	State string     `json:"state"`
	User  GithubUser `json:"user"`
}

type PullRequestInformation struct {
	Url    string             `json:"html_url"`
	Id     int                `json:"id"`
	Number int                `json:"number"`
	Title  string             `json:"title"`
	User   GithubUser         `json:"user"`
	Labels []PullRequestLabel `json:"labels"`
}

type PullRequestLabel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
