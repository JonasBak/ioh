package mqtt

import (
  "fmt"
  "strings"
  "os"
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "github.com/JonasBak/ioh/hub/ioh_config"
)

func discoverHandler(client MQTT.Client, msg MQTT.Message) {
  host := strings.Split(msg.Topic(), "/")[2]
  fmt.Printf("recieved message: %s, from client %s\n", msg.Payload(), host)

  payload := string(msg.Payload())

  if payload == TYPE_DISCOVER_EMPTY {
    config := ioh_config.GetConfig()
    requested_config := config.GetConfig(host)

    client.Publish(get_topic_client_hub(host), 0, false, TYPE_DISCOVER_ACK)
    if requested_config == nil {
      config.AddClient(host)
    } else {
      str := requested_config.ToString()
      client.Publish(get_topic_client_config(host), 0, false, str)
    }
  }
}

func defaultHandler(client MQTT.Client, msg MQTT.Message) {
  fmt.Printf("recieved message: %s\n", msg.Payload())
}

func ConnectAndListen() {
  opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt_broker:1883")

  opts.SetClientID(fmt.Sprintf("ioh-hub-sub-%s", os.Getenv("HOSTNAME")))
  opts.SetDefaultPublishHandler(defaultHandler)

  opts.OnConnect = func(c MQTT.Client) {
    if token := c.Subscribe(get_topic_client_discover("+"), 0, discoverHandler); token.Wait() && token.Error() != nil {
      panic(token.Error())
    }
  }
  client := MQTT.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  } else {
    fmt.Println("Connected to mqtt broker")
  }
}
