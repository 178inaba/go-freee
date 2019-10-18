package freee

import (
	"context"
	"net/http"
	"net/url"
)

// User is user.
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

// UserCompany is user company.
type UserCompany struct {
	ID            int    `json:"id"`
	DisplayName   string `json:"display_name"`
	Role          string `json:"role"`
	UseCustomRole bool   `json:"use_custom_role"`
}

// User returns login user.
func (c *Client) User(ctx context.Context) (*User, error) {
	var u struct {
		User `json:"user"`
	}

	q := url.Values{}
	q.Set("companies", "true")

	if err := c.do(ctx, http.MethodGet, "users/me", q, &u); err != nil {
		return nil, err
	}

	return &u.User, nil
}
