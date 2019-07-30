package mqtt

import (
  "fmt"
)

const (
  TYPE_DISCOVER_EMPTY = "EMPTY"
  TYPE_DISCOVER_ACK = "ACK"
  TYPE_STATUS_ON = "ON"
  TYPE_STATUS_OFF = "OFF"

  topic_client_discover = "ioh/client/%s/discover"
  topic_client_config = "ioh/client/%s/config"
  topic_client_hub = "ioh/client/%s/hub"
  topic_client_status = "ioh/client/%s/status"
)

func get_topic_client_discover(p string) string {
  return fmt.Sprintf(topic_client_discover, p)
}

func get_topic_client_config(p string) string {
  return fmt.Sprintf(topic_client_config, p)
}

func get_topic_client_hub(p string) string {
  return fmt.Sprintf(topic_client_hub, p)
}

func get_topic_client_status(p string) string {
  return fmt.Sprintf(topic_client_status, p)
}
