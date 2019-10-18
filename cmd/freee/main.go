package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/178inaba/go-freee"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Args[1]})
	tc := oauth2.NewClient(ctx, ts)

	c, err := freee.NewClient(tc)
	if err != nil {
		log.Fatal(err)
	}

	u, err := c.User(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", u)

	uc, err := c.UserCapability(ctx, 2000724)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", uc)
}
