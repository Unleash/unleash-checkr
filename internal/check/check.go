package check

import (
	"fmt"

	"github.com/apex/log"
	"github.com/gookit/color"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

// Start check
func Start(ctx *context.Context) error {
	log.WithFields(log.Fields{
		"url":           ctx.Config.URL,
		"expires":       ctx.Config.Expires,
		"slack-channel": ctx.Config.SlackChannel,
		"slack-token":   ctx.Config.SlackToken,
	}).Info("Start check")

	ff, err := flag.Get(ctx)
	if err != nil {
		return err
	}

	log.Info("Flags list")
	fmt.Println()

	for _, f := range ff {
		createdAt := color.New(color.Bold).Sprint("Created at:")

		log.Info(color.New(color.Bold, color.Green).Sprintf(f.Name))
		log.Info(f.Description)
		log.Info(fmt.Sprintf("%s %v", createdAt, f.CreatedAt))
		log.Info(fmt.Sprintf("%s/#/features/strategies/%s", ctx.Config.URL, f.Name))
		fmt.Println()
	}

	return nil
}
