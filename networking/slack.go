package networking

import (
	"github.com/Hurobaki/github-notifier/constants"
	"net/http"
)

func SendSlackMessage(req *http.Request) error {

	req.Header.Set(constants.HeaderContentType, "application/json")

	err := Post(req)

	if err != nil {
		return err
	}

	return nil
}
