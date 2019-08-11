package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var hubTopic = "ioh/client/tester/hub"
var configTopic = "ioh/client/tester/config"
var discoverTopic = "ioh/client/tester/discover"
var statusTopic = "ioh/client/tester/status"

func getTestClient(id string, hubChannel chan string, configChannel chan string) MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(os.Getenv("MQTT_BROKER"))

	opts.SetUsername(os.Getenv("MQTT_USER"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetClientID(fmt.Sprintf("tester-%s-%s", id, os.Getenv("HOSTNAME")))

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(hubTopic, 0,
			func(client MQTT.Client, msg MQTT.Message) {
				hubChannel <- string(msg.Payload())
			}); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		if token := c.Subscribe(configTopic, 0,
			func(client MQTT.Client, msg MQTT.Message) {
				configChannel <- string(msg.Payload())
			}); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func TestAck(t *testing.T) {
	hubChannel := make(chan string, 10)
	configChannel := make(chan string, 10)

	t.Log("Connecting to broker")
	client := getTestClient("ack", hubChannel, configChannel)

	t.Log("Publishing message to broker")
	client.Publish(discoverTopic, 0, false, "EMPTY")

	select {
	case res := <-hubChannel:
		t.Logf("Recieved: %s", res)
		if res != "ACK" {
			t.Error("Didn't recieve ack")
		}
	case <-time.After(10 * time.Second):
		t.Error("Timed out")
	}
}

var activeQuery = `{"query":"{\n  client(clientId: \"tester\"){\n    active\n    config {\n      plant\n      water\n    }\n  }  \n}","variables":null,"operationName":null}`

type ClientResponse struct {
	Data struct {
		Client struct {
			Active bool
			Config *struct {
				Plant string
				Water int
			}
		}
	}
}

func fetchClient() ClientResponse {
	resp, err := http.Post(fmt.Sprintf("%s/query", os.Getenv("HUB_URL")), "application/json", strings.NewReader(activeQuery))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	clientResponse := ClientResponse{}
	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		panic(err)
	}
	return clientResponse
}

func TestActiveAPI(t *testing.T) {
	hubChannel := make(chan string, 10)
	configChannel := make(chan string, 10)

	t.Log("Connecting to broker")
	client := getTestClient("active", hubChannel, configChannel)

	t.Log("Trying to deactivate")
	client.Publish(statusTopic, 0, false, "OFF")
	time.Sleep(1 * time.Second)
	if fetchClient().Data.Client.Active {
		t.Error("Client not deactivated")
	}

	t.Log("Trying to activate")
	client.Publish(statusTopic, 0, false, "ON")
	time.Sleep(1 * time.Second)
	if !fetchClient().Data.Client.Active {
		t.Error("Client not activated")
	}
}

var configQuery = `{"query":"mutation {\n  setConfig(config: {clientId: \"tester\", plant: \"test\", water: 6}) {\n    water\n    plant\n  }\n}\n","variables":null}`

func TestConfigAPI(t *testing.T) {
	hubChannel := make(chan string, 10)
	configChannel := make(chan string, 10)

	t.Log("Connecting to broker")
	client := getTestClient("config", hubChannel, configChannel)

	t.Log("Making shure client exists")
	client.Publish(discoverTopic, 0, false, "EMPTY")

	t.Log("Posting config")
	_, err := http.Post(fmt.Sprintf("%s/query", os.Getenv("HUB_URL")), "application/json", strings.NewReader(configQuery))
	if err != nil {
		panic(err)
	}

	t.Log("Checking mqtt")
	select {
	case res := <-configChannel:
		t.Logf("Recieved: %s", res)
		if res != "6" {
			t.Error("Didn't recieve correct config")
		}
	case <-time.After(10 * time.Second):
		t.Error("Timed out")
	}

	t.Log("Checking api")
	config := fetchClient().Data.Client.Config

	if config.Plant != "test" || config.Water != 6 {
		t.Error("Config not updated correctly")
	}
}
