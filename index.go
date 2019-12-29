package handler

import (
	"fmt"
	"github.com/Hurobaki/github-notifier/handlers"
	"github.com/Hurobaki/github-notifier/validator"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	if isPostMethod := validator.IsPostMethod(r); isPostMethod == false {
		fmt.Fprintf(w, "Not a valid method please send a POST method")
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	if isGithubWebHook := validator.IsGithubWebHook(r); isGithubWebHook == true {
		err := handlers.Github(w, r)

		if err == nil {
			w.WriteHeader(http.StatusOK)
		}
		return
	}

	http.Error(w, "Webhook source unknown", http.StatusForbidden)
	return
}
