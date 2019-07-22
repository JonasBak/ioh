package server

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/JonasBak/ioh/hub/ioh_config"
)

func ConfigHandler() Handler {
  config := ioh_config.GetConfig()
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

    var c ioh_config.PlantConfig
    defer r.Body.Close()
    json.NewDecoder(r.Body).Decode(&c)

    config.SetConfig(id[0], c)
  }
}
