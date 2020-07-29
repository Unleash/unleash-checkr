package slack

import (
	slackAPI "github.com/slack-go/slack"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

// Notify some slack channel
func Notify(ctx *context.Context) error {
	api := slackAPI.New(ctx.Config.SlackToken)

	var text = ""
	var iconURL = "https://raw.githubusercontent.com/Unleash/unleash-frontend/master/public/logo.png"
	_, _, _, err := api.SendMessage(ctx.Config.SlackChannel,
		slackAPI.MsgOptionAsUser(true),
		slackAPI.MsgOptionText(text, true),
		slackAPI.MsgOptionIconURL(iconURL),
		slackAPI.MsgOptionUsername("unleash-checkr"),
	)

	return err
}
