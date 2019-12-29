package validator

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/Hurobaki/github-notifier/constants"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func IsPostMethod(req *http.Request) bool {
	return req.Method == http.MethodPost
}

func IsGithubWebHook(req *http.Request) bool {
	validWebHook := regexp.MustCompile(constants.GithubHook)
	return validWebHook.MatchString(req.Header.Get(constants.HeaderUserAgent))
}

func ValidateSignature(req *http.Request) bool {
	switch req.Header.Get(constants.HeaderContentType) {
	case "application/json":
		const prefix string = "sha1="

		hash := hmac.New(sha1.New, []byte(os.Getenv("SECRET_TOKEN")))

		incomingHash := req.Header.Get(constants.GithubSecretHeader)

		body, err := ioutil.ReadAll(req.Body)

		req.Body.Close()
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err != nil {
			log.Printf("Cannot read the request body: %s\n", err)
			return false
		}

		hash.Write(body)
		expectedHash := prefix + hex.EncodeToString(hash.Sum(nil))

		return hmac.Equal([]byte(expectedHash), []byte(incomingHash))
	default:
		return false
	}
}
