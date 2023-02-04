package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type About struct {
	Name string `json:"name"`
}

func GetAbout(ctx echo.Context) error {
	about := About{
		Name: "Users",
	}
	ctx.Logger().Debug("About Request")
	return ctx.JSON(http.StatusOK, &about)
}
