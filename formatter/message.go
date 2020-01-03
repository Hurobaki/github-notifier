package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/Hurobaki/github-notifier/types"
)

func getSlackMessageTitle(status interface{}) string {
	switch status {
	case types.Opened:
		return "opened a new pull request"
	case types.Closed:
		return "closed a pull request"
	case types.ReviewRequest:
		return "requested a review on pull request"
	case types.Submitted:
		return "submitted a review on pull request"
	default:
		return "action undefined for the moment ..."
	}
}

//TODO format slack message for either pull_request or pull_request reviewer

func SlackMessage(status interface{}, user types.GithubUser, pr types.PullRequestInformation) ([]byte, error) {
	messageTitle := getSlackMessageTitle(status)

	message, err := json.Marshal(types.SlackBaseMessage{Text: fmt.Sprintf("%s %s", user.Login, messageTitle), Blocks: []interface{}{
		types.SlackContextElement{Type: "context", Elements: []interface{}{
			types.SlackImageElement{Type: "image", Image: pr.User.Avatar, Alt: "user image"},
			types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("<%s|%s> %s", user.Url, user.Login, messageTitle)},
		}},
		types.SlackSectionElement{Type: "section", Text: types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("*Pull request* \n *<%s|%s (#%d)>*", pr.Url, pr.Title, pr.Number)}},
		types.SlackDivider{Type: "divider"},
		types.SlackContextElement{Type: "context", Elements: []interface{}{
			types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("👀 Let's take a look !")},
		}},
	}})

	if err != nil {
		return nil, err
	}

	return message, nil
}
