package slack

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wesleimp/unleash-checkr/internal/flag"
	"github.com/wesleimp/unleash-checkr/pkg/config"
	"github.com/wesleimp/unleash-checkr/pkg/context"
)

func TestWithEmptyFlags(t *testing.T) {
	ctx := context.New(&config.Config{})
	var flags []flag.Flag
	text, err := describeMessageBody(ctx, flags)
	assert.NoError(t, err)
	assert.Equal(t, text, "Here are some flags that have passed their expiration and can now be removed.")
}

func TestWithOneFlag(t *testing.T) {
	ctx := context.New(&config.Config{
		URL: "http://test.com",
	})

	date, err := time.Parse(time.RFC3339, "2020-01-01T15:04:05+07:00")
	assert.NoError(t, err)

	var flags = []flag.Flag{
		{
			Name:        "First",
			Description: "",
			CreatedAt:   date,
		},
		{
			Name:        "Second",
			Description: "Second flag",
			CreatedAt:   date,
		},
	}
	text, err := describeMessageBody(ctx, flags)
	assert.NoError(t, err)
	assert.Equal(t, text, `Here are some flags that have passed their expiration and can now be removed.

*Flags*

*First*
Created at: 01/01/2020 3:04 PM
http://test.com/#/features/strategies/First

*Second* - Second flag
Created at: 01/01/2020 3:04 PM
http://test.com/#/features/strategies/Second
`)
}
