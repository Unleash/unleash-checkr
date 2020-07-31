package cli

import (
	"github.com/apex/log"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/internal/notification/slack"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

func runNotify(c *cli.Context) error {
	token := c.String("slack-token")
	channel := c.String("channel")
	url := c.String("url")
	expires := c.Int("expires")

	ctx := context.New(&config.Config{
		SlackChannel: channel,
		SlackToken:   token,
		URL:          url,
		Expires:      expires,
	})

	ff, err := flag.Get(ctx)
	if err != nil {
		return err
	}

	err = slack.Notify(ctx, ff)
	if err != nil {
		return err
	}

	log.Info("Flags sent to slack")

	return nil
}
