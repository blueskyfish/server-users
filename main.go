package main

import (
	"embed"
	"github.com/blueskyfish/server-users/configuration"
	"github.com/blueskyfish/server-users/http"
)

//go:embed web
var webFiles embed.FS

// main entry point of "Server Users"
func main() {
	conf := configuration.NewConfiguration()
	log := conf.WithGroup("main")

	server := http.NewServer(&conf)
	err := server.Init(&conf, &webFiles)
	if err != nil {
		panic(err)
	}

	log.Info("Server starts at ", server.Name)
	err = server.Start(&conf)
	if err != nil {
		panic(err)
	}
}
