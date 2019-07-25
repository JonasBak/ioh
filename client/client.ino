#include <ESP8266WiFi.h>
#include <PubSubClient.h>

WiFiClient espClient;
PubSubClient client(espClient);

const char* WIFI_SSID = "NextGenTel_E13E";
const char* WIFI_PASSWORD = "1C8CCDB197";

const char* MQTT_BROKER = "10.0.0.138";
const uint16_t MQTT_PORT = 1883;

char HOSTNAME[8];

void connect_wifi() {
  Serial.println("Connecting to WIFI...");
  WiFi.mode(WIFI_STA);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(500);
  }
  Serial.println("Connected!");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}

void connect_mqtt() {
  Serial.println("Connecting to MQTT...");
  
  client.setServer(MQTT_BROKER, MQTT_PORT);
  client.setCallback(callback);

  while (!client.connected()) {
    if (client.connect(HOSTNAME, NULL, NULL)) {
      Serial.println("Connected!");
    } else {
      Serial.print(".");
      delay(500);
    }
  }
  char topic[32];
  sprintf(topic, "ioh/client/%s/discover", HOSTNAME);
  client.subscribe(topic);
  char msg[128];
  sprintf(msg, "{\"ReqType\": \"discover_empty\", \"Host\": \"%s\"}", HOSTNAME);
  client.publish(topic, msg);
}

void setup() {
  Serial.begin(115200);
  String(ESP.getChipId(), HEX).toCharArray(HOSTNAME, 8);
  connect_wifi();
  connect_mqtt();
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.println((char *)payload);
}

void loop() {
  client.loop();
}
