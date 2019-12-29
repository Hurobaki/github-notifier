package decoder

import (
	"encoding/json"
	"github.com/Hurobaki/github-notifier/types"
	"net/http"
)

func GetPullRequest(req *http.Request) (types.PullRequest, error) {
	var pullRequest types.PullRequest

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&pullRequest)

	if err != nil {
		return pullRequest, err
	}

	return pullRequest, nil
}

func GetPullRequestReview(req *http.Request) (types.PullRequestReview, error) {
	var pullRequestReview types.PullRequestReview

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&pullRequestReview)

	if err != nil {
		return pullRequestReview, err
	}

	return pullRequestReview, nil
}