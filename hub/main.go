package main

import (
  "net/http"
  //"github.com/JonasBak/ioh/hub/ioh_config"
  "github.com/JonasBak/ioh/hub/server"
  "github.com/JonasBak/ioh/hub/mqtt"
)

func main() {
  // check flags for what to run

  go mqtt.ConnectAndListen()
  http.HandleFunc("/", server.ConfigHandler())
  http.ListenAndServe(":5151", nil)
}
