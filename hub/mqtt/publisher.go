package mqtt

import (
  "fmt"
  "os"
  "encoding/json"
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "github.com/JonasBak/ioh/hub/ioh_config"
)

type Publisher struct {
  c MQTT.Client
}

func GetPublisher() Publisher {
  opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt_broker:1883")

  opts.SetClientID(fmt.Sprintf("ioh-hub-pub-%s", os.Getenv("HOSTNAME")))
  opts.SetDefaultPublishHandler(defaultHandler)

  client := MQTT.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  } else {
    fmt.Println("Connected to mqtt broker")
  }

  return Publisher {client}
}

func (pub Publisher) UpdatedConfig(p string, c ioh_config.ClientConfig) {
  response := Req {
    ReqType: TYPE_CLIENT_UPDATED,
    Host: p,
    Config: &c,
  }
  str, err := json.Marshal(response)
  if err != nil {
    panic(err)
  }
  pub.c.Publish(get_topic_client_config(p), 0, false, str)
}
