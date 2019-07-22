package mqtt

import (
  "github.com/JonasBak/ioh/hub/ioh_config"
)

const (
  TYPE_DISCOVER_EMPTY = "discover_empty"
  TYPE_DISCOVER_ACK = "discover_ack"
  TYPE_DISCOVER_EXISTS = "discover_exists"
)

type Req struct {
  ReqType string
  Host string
  Config ioh_config.PlantConfig
}
