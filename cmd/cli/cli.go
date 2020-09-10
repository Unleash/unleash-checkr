package cli

import (
	"github.com/urfave/cli/v2"
)

// Run starts the app
func Run(version string, args []string) error {
	app := &cli.App{
		Name:    "unleash-checkr",
		Usage:   "Checks if any flags have expired and notifies somewhere",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "checks the flags and print on the console",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "url",
						Aliases:  []string{"u"},
						Usage:    "unleash api url",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "expires",
						Aliases: []string{"e"},
						Usage:   "expires after days",
						Value:   40,
					},
				},
				Action: runCheck,
			},
			{
				Name:    "notify",
				Aliases: []string{"n"},
				Usage:   "checks the flags and notify",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "channel",
						Usage:    "slack notification channel",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "slack-token",
						Usage:   "slack token",
						EnvVars: []string{"SLACK_TOKEN"},
					},
					&cli.StringFlag{
						Name:     "url",
						Aliases:  []string{"u"},
						Usage:    "unleash api url",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "expires",
						Aliases: []string{"e"},
						Usage:   "expires after days",
						Value:   40,
					},
				},
				Action: runNotify,
			},
		},
	}

	return app.Run(args)
}
