package cli

import (
	"github.com/apex/log"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/unleash-checkr/internal/check"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

func runCheck(c *cli.Context) error {
	expires := c.Int("expires")
	url := c.String("url")

	ctx := context.New(&config.Config{
		URL:     url,
		Expires: expires,
	})

	_, err := check.Start(ctx)
	if err != nil {
		log.WithError(err).Error("Error checking flags")
		return err
	}

	return nil
}
