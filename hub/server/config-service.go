package server

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/JonasBak/ioh/hub/ioh_config"
  "github.com/JonasBak/ioh/hub/mqtt"
)

func ConfigHandler() Handler {
  config := ioh_config.GetConfig()
  publisher := mqtt.GetPublisher()
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
    if r.Method != http.MethodPost {
      str, err := json.Marshal(config.GetUnconfigured())
      if err != nil {
        panic(err)
      }
      fmt.Fprintf(w, string(str))
    }

    id, ok := r.URL.Query()["id"]
    if !ok || len(id) != 1 {
      return
    }

    var c ioh_config.ClientConfig
    defer r.Body.Close()
    json.NewDecoder(r.Body).Decode(&c)

    config.SetConfig(id[0], c)

    publisher.UpdatedConfig(id[0], c)
  }
}
