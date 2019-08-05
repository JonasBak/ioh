package mqtt

import (
	"fmt"
	"github.com/JonasBak/ioh/hub/ioh_config"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type Publisher struct {
	c MQTT.Client
}

func GetPublisher() Publisher {
	opts := MQTT.NewClientOptions().AddBroker(os.Getenv("MQTT_BROKER"))

	opts.SetUsername(os.Getenv("MQTT_USER"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetClientID(fmt.Sprintf("ioh-hub-pub-%s", os.Getenv("HOSTNAME")))
	opts.SetDefaultPublishHandler(defaultHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Println("Connected to mqtt broker")
	}

	return Publisher{client}
}

func (pub Publisher) UpdatedConfig(p string, c ioh_config.ClientConfig) {
	str := c.ToString()
	pub.c.Publish(get_topic_client_config(p), 0, false, str)
}
