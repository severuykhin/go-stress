package client

import (
	"io"
	"net/http"
)

type httpClient struct {
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

func (ht *httpClient) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
func (ht *httpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return http.Post(url, contentType, body)
}
