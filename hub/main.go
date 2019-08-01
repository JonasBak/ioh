package main

import (
	"github.com/JonasBak/ioh/hub/ioh_config"
	"github.com/JonasBak/ioh/hub/mqtt"
	"github.com/JonasBak/ioh/hub/server"
	"net/http"
)

func main() {
	// check flags for what to run

	go mqtt.ConnectAndListen()

	config := ioh_config.GetConfig()
	publisher := mqtt.GetPublisher()
	http.HandleFunc("/", server.GQLHandler(config, publisher))
	http.HandleFunc("/query", server.QueryHandler(config, publisher))
	http.ListenAndServe(":5151", nil)
}
