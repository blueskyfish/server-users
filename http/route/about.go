package route

import (
	"github.com/blueskyfish/server-users/http/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type About struct {
	Name string `json:"name"`
}

func GetAbout(ctx echo.Context) error {
	serverCtx := context.ToServerContext(ctx)

	about := About{
		Name: "Users",
	}

	serverCtx.Logger().Debugf("About Request (has repository = %t)", serverCtx.HasRepository())
	return ctx.JSON(http.StatusOK, &about)
}
