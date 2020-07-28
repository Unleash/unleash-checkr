package check

import (
	"github.com/apex/log"
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

	return nil
}
