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
    str, _ := json.Marshal(config.GetUnconfigured())
    fmt.Fprintf(w, string(str))
  }
}
