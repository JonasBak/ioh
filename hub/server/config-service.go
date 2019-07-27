package server

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/JonasBak/ioh/hub/ioh_config"
  "github.com/JonasBak/ioh/hub/mqtt"
)

func ConfigHandler(config ioh_config.IOHConfig, publisher mqtt.Publisher) Handler {
  return func(w http.ResponseWriter, r *http.Request) {
    id, ok := r.URL.Query()["id"]
    if !ok || len(id) != 1 {
      return
    }

    if r.Method != http.MethodPost {
      str, err := json.Marshal(config.GetConfig(id[0]))
      if err != nil {
        panic(err)
      }
      fmt.Fprintf(w, string(str))
      return
    }

    var c ioh_config.ClientConfig
    defer r.Body.Close()
    json.NewDecoder(r.Body).Decode(&c)

    config.SetConfig(id[0], c)

    publisher.UpdatedConfig(id[0], c)
  }
}

func UnconfiguredHandler(config ioh_config.IOHConfig, publisher mqtt.Publisher) Handler {
  return func(w http.ResponseWriter, r *http.Request) {
    str, err := json.Marshal(config.GetUnconfigured())
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w, string(str))
  }
}

func ConfiguredHandler(config ioh_config.IOHConfig, publisher mqtt.Publisher) Handler {
  return func(w http.ResponseWriter, r *http.Request) {
    str, err := json.Marshal(config.GetConfigured())
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w, string(str))
  }
}
