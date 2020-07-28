package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleimp/unleash-checkr/pkg/config"
)

func TestNew(t *testing.T) {
	var ctx = New(&config.Config{
		URL: "http://hey.com",
	})

	assert.Equal(t, "http://hey.com", ctx.Config.URL)
}
