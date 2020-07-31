package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

func TestInvalidTemplate(t *testing.T) {
	ctx := context.New(&config.Config{})
	_, err := New(ctx).Apply("{{{.Foo}")
	assert.EqualError(t, err, "template: tmpl:1: unexpected \"{\" in command")
}

func TestWithExtraFields(t *testing.T) {
	ctx := context.New(&config.Config{})
	out, err := New(ctx).WithExtraFields(Fields{
		"Foo": "Bar",
	}).Apply("{{.Foo}}")
	assert.NoError(t, err)
	assert.Equal(t, out, "Bar")
}
