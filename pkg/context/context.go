package context

import (
	ctx "context"

	"github.com/wesleimp/unleash-checkr/pkg/config"
)

// Context struct
type Context struct {
	ctx.Context
	Config *config.Config
}

// New creates new context
func New(conf *config.Config) *Context {
	return Wrap(ctx.Background(), conf)
}

// Wrap context
func Wrap(ctx ctx.Context, conf *config.Config) *Context {
	return &Context{
		Context: ctx,
		Config:  conf,
	}
}
