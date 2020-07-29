package check

import (
	"fmt"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

// Start check
func Start(ctx *context.Context) ([]flag.Flag, error) {
	log.WithFields(log.Fields{
		"url":     ctx.Config.URL,
		"expires": ctx.Config.Expires,
	}).Info("Start check")

	ff, err := flag.Get(ctx)
	if err != nil {
		return nil, err
	}

	log.Info("Flags list")
	fmt.Println()

	var bold = color.New(color.Bold)
	for _, f := range ff {
		createdAt := bold.Sprint("Created at:")

		log.Info(color.New(color.Bold, color.FgGreen).Sprintf(f.Name))
		log.Info(f.Description)
		log.Info(fmt.Sprintf("%s %v", createdAt, f.CreatedAt))
		log.Info(fmt.Sprintf("%s/#/features/strategies/%s", ctx.Config.URL, f.Name))
		fmt.Println()
	}

	log.Info(fmt.Sprintf("%s %v", bold.Sprintf("Count:"), len(ff)))
	fmt.Println()

	return ff, nil
}
