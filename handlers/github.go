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
	"net/http"
	"os"
)

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
			fmt.Fprintf(w, "Problem with decoder %s ", err)
			return err
		}
		message, err = formatter.SlackMessage(pr.Action, pr.Information)

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
		message, err = formatter.SlackMessage(pr.Action, pr.Information)
	}

	if err != nil {
		fmt.Fprintf(w, "Problem with message formatter %s ", err)
		return err
	}
	request, err := networking.HttpRequest(http.MethodPost, os.Getenv("SLACK_URL"), message)

	if err != nil {
		fmt.Fprintf(w, "Problem with request builder %s ", err)
		return err
	}

	err = networking.SendSlackMessage(request)

	if err != nil {
		fmt.Fprintf(w, "Problem sending slack message %s", err)
		return err
	}

	return nil
}
