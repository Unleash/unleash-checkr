package tmpl

import (
	"bytes"
	"html/template"

	"github.com/wesleimp/unleash-checkr/pkg/context"
)

// Template holds data that can be applied to a template string
type Template struct {
	fields Fields
}

// Fields that will be available to the template engine
type Fields map[string]interface{}

var (
	url     = "URL"
	expires = "Expires"
)

// New Template
func New(ctx *context.Context) *Template {
	return &Template{
		fields: Fields{
			url:     ctx.Config.URL,
			expires: ctx.Config.Expires,
		},
	}
}

// WithExtraFields allows to add new more custom fields to the template
func (t *Template) WithExtraFields(f Fields) *Template {
	for k, v := range f {
		t.fields[k] = v
	}

	return t
}

// Apply applies the given string against the Fields stored in the template
func (t *Template) Apply(s string) (string, error) {
	var output bytes.Buffer
	tmpl, err := template.New("tmpl").Option("missingkey=error").Parse(s)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&output, t.fields)
	return output.String(), err
}
