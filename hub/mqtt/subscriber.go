package mqtt

import (
  "fmt"
  "os"
  "encoding/json"
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "github.com/JonasBak/ioh/hub/ioh_config"
)

func hubHandler(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("recieved message: %s\n", msg.Payload())
}

func discoverHandler(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("recieved message: %s\n", msg.Payload())

    var req Req
    json.Unmarshal(msg.Payload(), &req)
    config := ioh_config.GetConfig()

    if req.ReqType == TYPE_DISCOVER_EMPTY {
      requested_config := config.GetConfig(req.Host)
      var response Req
      if requested_config == nil {
        config.AddUnconfigured(req.Host)
        response = Req {
          ReqType: TYPE_DISCOVER_ACK,
          Host: req.Host,
        }
      } else {
        response = Req {
          ReqType: TYPE_DISCOVER_EXISTS,
          Host: req.Host,
          Config: *requested_config,
        }
      }
      str, err := json.Marshal(response)
      if err != nil {
        panic(err)
      }
      client.Publish(msg.Topic(), 0, false, str)
    }
}

func defaultHandler(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("recieved message: %s\n", msg.Payload())
}

func ConnectAndListen() {
  opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt_broker:1883")

  opts.SetClientID(fmt.Sprintf("ioh-hub-%s", os.Getenv("HOSTNAME")))
  opts.SetDefaultPublishHandler(defaultHandler)

  topicHub := "ioh/hub"
  topicDiscover := "ioh/discover"

  opts.OnConnect = func(c MQTT.Client) {
    if token := c.Subscribe(topicHub, 0, hubHandler); token.Wait() && token.Error() != nil {
      panic(token.Error())
    }
    if token := c.Subscribe(topicDiscover, 0, discoverHandler); token.Wait() && token.Error() != nil {
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
