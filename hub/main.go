package main

import (
  "net/http"
  "github.com/JonasBak/ioh/hub/ioh_config"
  "github.com/JonasBak/ioh/hub/server"
  "github.com/JonasBak/ioh/hub/mqtt"
)

func main() {
  // check flags for what to run

  go mqtt.ConnectAndListen()

  config := ioh_config.GetConfig()
  publisher := mqtt.GetPublisher()
  http.HandleFunc("/config", server.ConfigHandler(config, publisher))
  http.HandleFunc("/unconfigured", server.UnconfiguredHandler(config, publisher))
  http.HandleFunc("/configured", server.ConfiguredHandler(config, publisher))
  http.ListenAndServe(":5151", nil)
}
