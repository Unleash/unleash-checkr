package slack

import (
	"fmt"

	slackAPI "github.com/slack-go/slack"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

var (
	message = "Here are some flags that have passed their expiration and can now be removed. :rocket:"
	color   = "#36a64f"
)

// Notify some slack channel
func Notify(ctx *context.Context, flags []flag.Flag) error {
	api := slackAPI.New(ctx.Config.SlackToken)

	att := createAttachments(ctx, flags)

	var iconURL = "https://raw.githubusercontent.com/Unleash/unleash-frontend/master/public/logo.png"
	_, _, _, err := api.SendMessage(ctx.Config.SlackChannel,
		slackAPI.MsgOptionText(message, true),
		slackAPI.MsgOptionIconURL(iconURL),
		slackAPI.MsgOptionUsername("unleash-checkr"),
		slackAPI.MsgOptionAttachments(att...),
	)

	return err
}

func createAttachments(ctx *context.Context, flags []flag.Flag) []slackAPI.Attachment {
	var aa []slackAPI.Attachment

	for _, f := range flags {
		aa = append(aa, slackAPI.Attachment{
			Color:     color,
			Title:     f.Name,
			TitleLink: fmt.Sprintf("%s/#/features/strategies/%s", ctx.Config.URL, f.Name),
			Text:      f.Description,
			Fields: []slackAPI.AttachmentField{
				{
					Title: "Created at",
					Value: f.CreatedAt.Format("01/02/2006 3:04 PM"),
				},
			},
		})
	}

	return aa
}
