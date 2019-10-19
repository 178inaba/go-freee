package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/178inaba/go-freee"
	"github.com/urfave/cli"
	"golang.org/x/xerrors"
)

func user(c *cli.Context) error {
	ctx := c.App.Metadata["context"].(context.Context)
	fc := c.App.Metadata["freee_client"].(*freee.Client)

	u, err := fc.User(ctx)
	if err != nil {
		return xerrors.Errorf("user: %w", err)
	}

	// TODO
	fmt.Printf("%#v\n", u)

	return nil
}

func userCapability(c *cli.Context) error {
	ctx := c.App.Metadata["context"].(context.Context)
	fc := c.App.Metadata["freee_client"].(*freee.Client)

	args := c.Args()
	if !args.Present() {
		return errors.New("argument is required")
	}

	companyIDStr := args.First()
	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		return xerrors.Errorf("invalid company ID: %q: %w", companyIDStr, err)
	}

	uc, err := fc.UserCapability(ctx, companyID)
	if err != nil {
		return xerrors.Errorf("user capability: %w", err)
	}

	// TODO
	fmt.Printf("%#v\n", uc)

	return nil
}
