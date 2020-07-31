package slack

import (
	"fmt"

	slackAPI "github.com/slack-go/slack"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/internal/tmpl"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

var template = `Here are some flags that have passed their expiration and can now be removed.

{{- with .Flags }}

*Flags*
{{ range $element := .}}
*{{ .Name }}*{{ if .Description }} - {{ .Description }}{{ end }}
Created at: {{ .CreatedAt }}
{{ .Location }}
{{ end -}}
{{- end -}}
`

// Notify some slack channel
func Notify(ctx *context.Context, flags []flag.Flag) error {
	api := slackAPI.New(ctx.Config.SlackToken)

	text, err := describeMessageBody(ctx, flags)
	if err != nil {
		return err
	}

	var iconURL = "https://raw.githubusercontent.com/Unleash/unleash-frontend/master/public/logo.png"
	_, _, _, err = api.SendMessage(ctx.Config.SlackChannel,
		slackAPI.MsgOptionAsUser(true),
		slackAPI.MsgOptionText(text, true),
		slackAPI.MsgOptionIconURL(iconURL),
		slackAPI.MsgOptionUsername("unleash-checkr"),
	)

	return err
}

func describeMessageBody(ctx *context.Context, flags []flag.Flag) (string, error) {
	var ff = parseToTemplate(ctx, flags)

	text, err := tmpl.New(ctx).
		WithExtraFields(tmpl.Fields{
			"Flags": ff,
		}).
		Apply(template)
	if err != nil {
		return "", err
	}

	return text, nil
}

func parseToTemplate(ctx *context.Context, flags []flag.Flag) []map[string]string {
	var ff []map[string]string
	for _, f := range flags {
		ff = append(ff, map[string]string{
			"Name":        f.Name,
			"Description": f.Description,
			"CreatedAt":   f.CreatedAt.Format("01/02/2006 3:04 PM"),
			"Location":    fmt.Sprintf("%s/#/features/strategies/%s", ctx.Config.URL, f.Name),
		})
	}

	return ff
}
