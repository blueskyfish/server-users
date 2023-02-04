package http

import (
	"fmt"
	"github.com/blueskyfish/server-users/configuration"
	"github.com/blueskyfish/server-users/http/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LogMiddleware set the own logger
func LogMiddleware(conf *configuration.Configuration, group string) echo.MiddlewareFunc {
	// create the logger with
	logger := log.NewHttpLogger(conf, group)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.SetLogger(logger)
			return next(ctx)
		}
	}
}

func RequestMiddleware(conf *configuration.Configuration) echo.MiddlewareFunc {
	requestLogger := conf.WithGroup("request")
	logUserAgent := conf.Stage() == configuration.StateProd
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogLatency:   true,
		LogUserAgent: logUserAgent,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger := requestLogger
			if logUserAgent {
				logger = logger.WithField("agent", v.UserAgent)
			}
			logger.Info(fmt.Sprintf("method=%s uri=%s status=%d duration=%s", v.Method, v.URI, v.Status, v.Latency))
			return nil
		},
	})
}
