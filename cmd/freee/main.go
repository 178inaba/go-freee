package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/178inaba/go-freee"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

const loginCommandName = "login"

func main() {
	os.Exit(run())
}

func run() int {
	app := cli.NewApp()
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Metadata = map[string]interface{}{"context": context.Background()}
	app.Before = before
	app.Commands = []cli.Command{{
		Name:   loginCommandName,
		Usage:  "Login freee account.",
		Action: login,
	}, {
		Name:   "user",
		Usage:  "Get login user.",
		Action: user,
	}, {
		Name:   "user-capability",
		Usage:  "Get login user capability.",
		Action: userCapability,
	}}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(app.ErrWriter, "freee: %v.\n", err)
		return 1
	}

	return 0
}

func before(c *cli.Context) error {
	ctx := c.App.Metadata["context"].(context.Context)

	oc := oauth2.Config{
		ClientID:     "fd6ba88f104d8cb7b7a394cfd47845e5ea93157d1185e6e6cf5fb5747a085031",
		ClientSecret: "b8a1042992a87466b4e37ce8a35126001b75d3a3f733ce595ac91429619aff90",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.secure.freee.co.jp/public_api/authorize",
			TokenURL: "https://accounts.secure.freee.co.jp/public_api/token",
		},
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
	}

	d, err := os.UserHomeDir()
	if err != nil {
		return xerrors.Errorf("user home directory: %w", err)
	}
	fp := filepath.Join(d, ".freee", "credentials.json")

	var freeeClient *freee.Client
	if c.Args().First() != loginCommandName {
		f, err := os.Open(fp)
		if err != nil {
			return xerrors.Errorf("open credentials file: %w", err)
		}
		defer f.Close()

		var t oauth2.Token
		if err := json.NewDecoder(f).Decode(&t); err != nil {
			return xerrors.Errorf("decode credentials json: %w", err)
		}

		fc, err := freee.NewClient(oc.Client(ctx, &t))
		if err != nil {
			return xerrors.Errorf("new freee client: %w", err)
		}
		freeeClient = fc
	}

	c.App.Metadata["oauth2_client"] = &oc
	c.App.Metadata["credentials_filepath"] = fp
	c.App.Metadata["freee_client"] = freeeClient

	return nil
}

func login(c *cli.Context) error {
	ctx := c.App.Metadata["context"].(context.Context)
	oc := c.App.Metadata["oauth2_client"].(*oauth2.Config)
	fp := c.App.Metadata["credentials_filepath"].(string)

	fmt.Fprint(c.App.Writer, "Go to the following link in your browser:\n\n")
	fmt.Fprintf(c.App.Writer, "    %s\n\n\n", oc.AuthCodeURL(""))
	fmt.Fprint(c.App.Writer, "Enter verification code: ")

	var code string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		code = s.Text()
		break
	}
	if err := s.Err(); err != nil {
		return xerrors.Errorf("scan code: %w", err)
	}

	t, err := oc.Exchange(ctx, code)
	if err != nil {
		return xerrors.Errorf("exchange oauth2 code to token: %w", err)
	}

	if err := os.Mkdir(filepath.Dir(fp), 0700); err != nil && !os.IsExist(err) {
		return xerrors.Errorf("create credentials directory: %w", err)
	}

	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("create credentials file: %w", err)
	}
	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	if err := e.Encode(t); err != nil {
		return xerrors.Errorf("encode json: %w", err)
	}

	fc, err := freee.NewClient(oc.Client(ctx, t))
	if err != nil {
		return xerrors.Errorf("new freee client: %w", err)
	}

	u, err := fc.User(ctx)
	if err != nil {
		return xerrors.Errorf("user: %w", err)
	}

	fmt.Fprintf(c.App.Writer, "You are now logged in as [%s].\n", u.Email)

	return nil
}
