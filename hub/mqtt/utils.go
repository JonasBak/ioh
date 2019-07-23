package mqtt

import (
  "fmt"
  "github.com/JonasBak/ioh/hub/ioh_config"
)

const (
  TYPE_DISCOVER_EMPTY = "discover_empty"
  TYPE_DISCOVER_ACK = "discover_ack"
  TYPE_DISCOVER_EXISTS = "discover_exists"

  TYPE_CLIENT_UPDATED = "client_updated"

  topic_client_discover = "ioh/client/%s/discover"
  topic_client_config = "ioh/client/%s/config"
)

func get_topic_client_discover(p string) string {
  return fmt.Sprintf(topic_client_discover, p)
}

func get_topic_client_config(p string) string {
  return fmt.Sprintf(topic_client_config, p)
}

type Req struct {
  ReqType string
  Host string
  Config *ioh_config.ClientConfig
}
