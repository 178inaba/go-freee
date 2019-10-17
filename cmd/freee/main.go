package main

import (
	"context"

	"github.com/178inaba/go-freee"
	"golang.org/x/oauth2"
)

func main() {
	// TODO
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "... your access token ..."})
	tc := oauth2.NewClient(ctx, ts)

	freee.NewClient(tc)
}