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
				Name:      "check",
				Aliases:   []string{"c"},
				Usage:     "checks the flags and print on the console",
				ArgsUsage: "URL",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "expires",
						Aliases: []string{"e"},
						Usage:   "expires after days",
						Value:   7,
					},
				},
				Action: runCheck,
			},
		},
	}

	return app.Run(args)
}
