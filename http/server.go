package http

import (
	"embed"
	"github.com/blueskyfish/server-users/configuration"
	"github.com/blueskyfish/server-users/http/log"
	"github.com/blueskyfish/server-users/route"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"io/fs"
	"net/http"
)

// Server the server struct
type Server struct {
	// Name the server name
	Name string

	http *echo.Echo
}

func NewServer(conf *configuration.Configuration) Server {
	server := Server{}
	server.Name = conf.Name()
	server.http = echo.New()
	server.http.Logger = log.NewHttpLogger(conf, "server")
	server.http.HideBanner = true
	return server
}

// Init initialize the server middlewares, routes and embed resources
func (server Server) Init(conf *configuration.Configuration, webFiles *embed.FS) error {
	serverLogger := conf.WithGroup("server")

	// Middlewares
	server.http.Use(LogMiddleware(conf, "context"))
	server.http.Use(RequestMiddleware(conf))
	server.http.Use(mw.Recover())
	serverLogger.Debug("add middlewares")

	// Routes
	server.http.GET("/api/about", route.GetAbout)
	serverLogger.Debug("add routes")

	// Embed resources
	fileSys, err := fs.Sub(webFiles, "web")
	if err != nil {
		return err
	}
	assetHandler := http.FileServer(http.FS(fileSys))
	server.http.GET("/*", echo.WrapHandler(assetHandler))
	serverLogger.Debug("setup asset handler")
	return nil
}

// Start starts the http server. Before starting, it is connected with the embed resources
func (server Server) Start(conf *configuration.Configuration) error {
	return server.http.Start(conf.Address())
}

// Http returns the http instance
func (server Server) Http() *echo.Echo {
	return server.http
}
