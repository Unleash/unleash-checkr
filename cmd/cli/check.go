package cli

import (
	"errors"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/unleash-checkr/internal/check"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

func runCheck(c *cli.Context) error {
	expires := c.Int("expires")

	url, err := checkArgs(c.Args())
	if err != nil {
		return err
	}

	ctx := context.New(&config.Config{
		URL:     url,
		Expires: expires,
	})

	err = check.Start(ctx)
	if err != nil {
		log.WithError(err).Error("Error checking flags")
		return err
	}

	return nil
}

func checkArgs(args cli.Args) (string, error) {
	if args.Len() == 0 {
		return "", errors.New("check command requires exactly 1 argument. See --help")
	}

	return args.Get(0), nil
}
