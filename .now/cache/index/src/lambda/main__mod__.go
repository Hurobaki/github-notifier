package main

import (
  "net/http"
  "github-notifier/api"

  now "github.com/zeit/now/utils/go/bridge"
)

func main() {
  now.Start(http.HandlerFunc(api.Handler))
}
