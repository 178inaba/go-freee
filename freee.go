package freee

import (
	"net/http"
	"net/url"
)

const defaultBaseRawurl = "https://api.freee.co.jp/"

// Client is freee client.
type Client struct {
	BaseURL *url.URL

	client *http.Client
}

// NewClient returns freee client.
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	u, err := url.Parse(defaultBaseRawurl)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:  httpClient,
		BaseURL: u,
	}, nil
}
