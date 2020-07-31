package slack

import (
	"fmt"

	"github.com/apex/log"
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
	_, _, _, err := api.SendMessage(ctx.Config.SlackChannel,
		slackAPI.MsgOptionText(message, true),
		slackAPI.MsgOptionUsername("unleash-checkr"),
		slackAPI.MsgOptionAttachments(att...),
	)

	return err
}

func createAttachments(ctx *context.Context, flags []flag.Flag) []slackAPI.Attachment {
	log.Infof("Creating %v attachments.", len(flags))

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
			Footer:     "unleash",
			FooterIcon: "https://emoji.slack-edge.com/TGN5JSCV7/unleash/e4e224dacc909a28.png",
		})
	}

	return aa
}
