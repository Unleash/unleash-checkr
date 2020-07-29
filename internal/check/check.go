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
		name := color.New(color.Bold).Sprintf("Name:")
		descripion := color.New(color.Bold).Sprint("Description:")
		createdAt := color.New(color.Bold).Sprint("Created at:")

		log.Info(fmt.Sprintf("%s %s", name, f.Name))
		log.Info(fmt.Sprintf("%s %s", descripion, f.Description))
		log.Info(fmt.Sprintf("%s %v", createdAt, f.CreatedAt))
		fmt.Println()
	}

	return nil
}
