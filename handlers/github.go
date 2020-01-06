package handlers

import (
	"errors"
	"fmt"
	"github.com/Hurobaki/github-notifier/constants"
	"github.com/Hurobaki/github-notifier/decoder"
	"github.com/Hurobaki/github-notifier/formatter"
	"github.com/Hurobaki/github-notifier/networking"
	"github.com/Hurobaki/github-notifier/types"
	"github.com/Hurobaki/github-notifier/validator"
	"log"
	"net/http"
	"os"
)

var validPullRequestStatus = [...]types.PullRequestStatus{
	types.Opened, types.Closed, types.Labeled, types.ReviewRequest,
}

func isValidPullRequestStatus(pr types.PullRequestStatus) bool {
	for _, status := range validPullRequestStatus {
		fmt.Println(status)
		fmt.Println(pr)
		if status == pr {
			return true
		}
	}
	return false
}

func Github(w http.ResponseWriter, r *http.Request) error {
	if isValid := validator.ValidateSignature(r); isValid == false {
		fmt.Println("Wrong secret token provided")
		http.Error(w, "Wrong secret token provided", http.StatusForbidden)
		return errors.New("wrong secret token provided")
	}
	var err error
	var message []byte
	eventName := r.Header.Get(constants.GithubEvent)

	// TODO: find a way to refacto this code

	switch eventName {
	case types.PullRequestEvent:
		pr, err := decoder.GetPullRequest(r)
		if err != nil {
			fmt.Println("Problem with decoder :", err)
			fmt.Fprintf(w, "Problem with decoder %s ", err)
			return err
		}

		if isValidStatus := isValidPullRequestStatus(pr.Action); isValidStatus == true {
			message, err = formatter.SlackMessage(pr.Action, pr.Information.User, pr.Information, pr.Sender)
		} else {
			log.Println("Pull request status unsupported ", pr.Action)
			return errors.New("pull request status unsupported")
		}

	case types.PullRequestReviewEvent:
		pr, err := decoder.GetPullRequestReview(r)
		if err != nil {
			fmt.Fprintf(w, "Problem with decoder %s ", err)
			return err
		}
		// Handle only submitted Pull Request Review
		if pr.Action != types.Submitted {
			return nil
		}
		message, err = formatter.SlackMessage(pr.Action, pr.Review.User, pr.Information, pr.Sender)
	}

	if err != nil {
		log.Println("Problem with message formatter ", err)
		fmt.Fprintf(w, "Problem with message formatter %s ", err)
		return err
	}

	log.Println("SLACK URL", os.Getenv("SLACK_URL"))
	request, err := networking.HttpRequest(http.MethodPost, os.Getenv("SLACK_URL"), message)

	if err != nil {
		log.Println("Problem with request builder", err)
		fmt.Fprintf(w, "Problem with request builder %s ", err)
		return err
	}

	err = networking.SendSlackMessage(request)

	if err != nil {
		log.Println("Problem sending slack message ", err)
		fmt.Fprintf(w, "Problem sending slack message %s", err)
		return err
	}

	return nil
}
