package utils

import (
	"net/http"
	"net/url"
)

func HTTPGet(url string, query *url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return http.DefaultClient.Do(req)
}
