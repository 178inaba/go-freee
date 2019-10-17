package freee

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
)

type User struct {
	ID            int           `json:"id"`
	Email         string        `json:"email"`
	DisplayName   string        `json:"display_name"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	FirstNameKana string        `json:"first_name_kana"`
	LastNameKana  string        `json:"last_name_kana"`
	Companies     []UserCompany `json:"companies"`
}

type UserCompany struct {
	ID            int    `json:"id"`
	DisplayName   string `json:"display_name"`
	Role          string `json:"role"`
	UseCustomRole bool   `json:"use_custom_role"`
}

// User returns login user.
func (c *Client) User(ctx context.Context) (*User, error) {
	u, err := c.BaseURL.Parse(path.Join(c.BaseURL.Path, "users/me"))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var uw struct {
		User `json:"user"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&uw); err != nil {
		return nil, err
	}

	return &uw.User, nil
}
