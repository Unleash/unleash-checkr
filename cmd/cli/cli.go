package cli

import (
	"fmt"

	"github.com/apex/log"
	clih "github.com/apex/log/handlers/cli"
	"github.com/urfave/cli/v2"

	"github.com/wesleimp/unleash-checkr/internal/check"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

// Run starts the app
func Run(version string, args []string) error {
	app := &cli.App{
		Name:    "unleash-checkr",
		Usage:   "Checks if any flags have expired and notifies somewhere",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "unleash url",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "expires",
				Aliases: []string{"e"},
				Usage:   "expires after days",
				Value:   7,
			},
			&cli.StringFlag{
				Name:  "slack-channel",
				Usage: "slack notification channel",
			},
			&cli.StringFlag{
				Name:    "slack-token",
				Usage:   "slack token",
				EnvVars: []string{"SLACK_TOKEN"},
			},
		},
		Action: run,
	}

	return app.Run(args)
}

func run(c *cli.Context) error {
	url := c.String("url")
	expires := c.Int("expires")
	slackChannel := c.String("slack-channel")
	slackToken := c.String("slack-token")

	log.SetHandler(clih.Default)

	fmt.Println()
	defer fmt.Println()

	ctx := context.New(&config.Config{
		URL:          url,
		Expires:      expires,
		SlackChannel: slackChannel,
		SlackToken:   slackToken,
	})

	err := check.Start(ctx)
	if err != nil {
		log.WithError(err).Error("Error checking flags")
		return err
	}

	return nil
}
