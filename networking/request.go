package networking

import (
	"bytes"
	"net/http"
)

func HttpRequest(method, url string, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	return request, nil
}
