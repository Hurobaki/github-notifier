package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PullRequest struct {
	Number string `json:"number"`
	Information PullRequestInformation `json:"pull_request"`
}

type PullRequestInformation struct {
	Url string
	Id int
}


func getPullRequestInformation(pr PullRequest, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&pr)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var pr PullRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pr)
	if err != nil {
		fmt.Fprintf(w, "error %s", err)
	}

	fmt.Fprintf(w, pr.Information.Url)

	url := "https://hooks.slack.com/services/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"

	var jsonStr = []byte(`{"text":"Buy cheese and bread for breakfast."}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	fmt.Fprintf(w, bodyString)

	fmt.Fprintf(w, "Index route called ")
}
