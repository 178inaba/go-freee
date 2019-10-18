package freee

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

const defaultBaseRawurl = "https://api.freee.co.jp/api/1/"

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

func (c *Client) do(ctx context.Context, method, apiPath string, query url.Values, v interface{}) error {
	u, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, apiPath))
	if err != nil {
		return err
	}
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
