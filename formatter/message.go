package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/Hurobaki/github-notifier/types"
	"log"
)

type NotificationTitle string

const (
	Opened        NotificationTitle = "opened a new pull request"
	Closed        NotificationTitle = "closed a pull request"
	ReviewRequest NotificationTitle = "requested a review on pull request"
	Submitted     NotificationTitle = "submitted a review on pull request"
	Labeled       NotificationTitle = "added a label on pull request"
	Unknown       NotificationTitle = "action undefined for the moment ..."
)

//TODO format slack message for either pull_request or pull_request reviewer

func SlackMessage(status interface{}, user types.GithubUser, pr types.PullRequestInformation, sender types.GithubUser) ([]byte, error) {
	var message types.SlackBaseMessage
	var blocks []interface{}
	var additionalInformation []interface{}

	notificationTitle := SlackNotificationTitle(sender)
	commonMessage := SlackMessageCommonBody(sender)

	switch status {
	case types.Opened:
		message.Text = notificationTitle(Opened)
		blocks = append(blocks, commonMessage(Opened))
	case types.Closed:
		message.Text = notificationTitle(Closed)
		blocks = append(blocks, commonMessage(Closed))
	case types.ReviewRequest:
		message.Text = notificationTitle(ReviewRequest)
		blocks = append(blocks, commonMessage(ReviewRequest))
	case types.Submitted:
		message.Text = notificationTitle(Submitted)
		blocks = append(blocks, commonMessage(Submitted))
	case types.Labeled:
		message.Text = notificationTitle(Labeled)
		blocks = append(blocks, commonMessage(Labeled))
		SlackMessageLabels(pr.Labels, &additionalInformation)
	default:
		message.Text = notificationTitle(Unknown)
		blocks = append(blocks, commonMessage(Unknown))
	}

	blocks = append(blocks, types.SlackSectionElement{Type: "section", Text: types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("*Pull request* \n *<%s|%s (#%d)>*", pr.Url, pr.Title, pr.Number)}})
	blocks = append(blocks, types.SlackSectionElement{Type: "section", Text: types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("*Author* \n *<%s|%s> created the pull request*", user.Url, user.Login)}})
	blocks = append(blocks, additionalInformation...)
	blocks = append(blocks, SlackMessageFooter()...)

	message.Blocks = blocks
	slackMessage, err := json.Marshal(message)

	if err != nil {
		log.Println("Error when Marshal slack message", err)
		return nil, err
	}

	return slackMessage, nil
}

func SlackNotificationTitle(user types.GithubUser) func(NotificationTitle) string {
	return func(notificationTitle NotificationTitle) string {
		return fmt.Sprintf("%s %s", user.Login, notificationTitle)
	}
}

func SlackMessageCommonBody(user types.GithubUser) func(title NotificationTitle) types.SlackContextElement {
	return func(title NotificationTitle) types.SlackContextElement {
		return types.SlackContextElement{Type: "context", Elements: []interface{}{
			types.SlackImageElement{Type: "image", Image: user.Avatar, Alt: "user image"},
			types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("<%s|%s> %s", user.Url, user.Login, title)},
		}}
	}
}

func SlackMessageFooter() []interface{} {
	return []interface{}{
		types.SlackDivider{Type: "divider"},
		types.SlackContextElement{Type: "context", Elements: []interface{}{
			types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("👀  Let's take a look !")},
		}},
	}
}

func SlackMessageLabels(labels []types.PullRequestLabel, additionalInformation *[]interface{}) {
	*additionalInformation = append(*additionalInformation, types.SlackSectionElement{Type: "section", Text: types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("*Labels*")}})
	labelsContainer := types.SlackContextElement{Type: "context"}
	var formattedLabels []interface{}

	for _, label := range labels {
		formattedLabels = append(formattedLabels, types.SlackTextElement{Type: "mrkdwn", Text: fmt.Sprintf("*%s*", label.Name)})
	}

	labelsContainer.Elements = formattedLabels
	*additionalInformation = append(*additionalInformation, labelsContainer)
}
