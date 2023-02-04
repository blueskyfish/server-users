package context

import (
	"github.com/blueskyfish/server-users/configuration"
	"github.com/blueskyfish/server-users/repository"
	"github.com/labstack/echo/v4"
)

// ServerContext extends the echo context and contains the property "Repository"
type ServerContext struct {
	echo.Context
	repository *repository.Repository
}

// HasRepository checks the existing repository
func (c ServerContext) HasRepository() bool {
	return c.repository != nil
}

func (c ServerContext) Repository() *repository.Repository {
	return c.repository
}

func NewServerContext(config *configuration.Configuration, ctx echo.Context) *ServerContext {
	c := ServerContext{Context: ctx}
	c.repository = config.Repository()
	return &c
}

// ToServerContext cast to the server context
func ToServerContext(ctx echo.Context) *ServerContext {
	return ctx.(*ServerContext)
}
