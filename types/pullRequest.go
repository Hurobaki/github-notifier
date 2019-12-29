package types

type PullRequestStatus string

type PullRequestReviewStatus string

const (
	Closed        PullRequestStatus = "closed"
	Opened        PullRequestStatus = "opened"
	ReviewRequest PullRequestStatus = "review_requested"
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
}

type PullRequest struct {
	Action      PullRequestStatus      `json:"action"`
	Number      string                 `json:"number"`
	Information PullRequestInformation `json:"pull_request"`
}

type PullRequestReviewInformation struct {
	State string `json:"state"`
}

type PullRequestInformation struct {
	Url    string     `json:"html_url"`
	Id     int        `json:"id"`
	Title  string     `json:"title"`
	Number int        `json:"number"`
	User   GithubUser `json:"user"`
}
