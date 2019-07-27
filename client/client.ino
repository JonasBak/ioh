#include <ESP8266WiFi.h>
#include <PubSubClient.h>

WiFiClient espClient;
PubSubClient client(espClient);

const char* WIFI_SSID = "NextGenTel_E13E";
const char* WIFI_PASSWORD = "1C8CCDB197";
const char* MQTT_BROKER = "10.0.0.138";
const int   MQTT_PORT = 1883;

char HOSTNAME[8];

enum STATE {
  EMPTY, ACKED, RUNNING
} current_state;

struct CONFIG {
  int water;
  unsigned long last_tick;
  bool pin_is_on;
} current_config;

const int LED_PIN = 2;
unsigned long LAST_CONNECT;

void connect_wifi() {
  Serial.print("Connecting to WIFI...");
  WiFi.mode(WIFI_STA);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  while (WiFi.status() != WL_CONNECTED) {
    Serial.print(".");
    delay(500);
  }
  Serial.println("\nConnected!");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}

void connect_mqtt() {
  Serial.print("Connecting to MQTT...");
  client.setServer(MQTT_BROKER, MQTT_PORT);
  client.setCallback(callback);
  while (!client.connected()) {
    if (!client.connect(HOSTNAME, NULL, NULL)) {
      Serial.print(".");
      delay(500);
    }
  }
  Serial.println("\nConnected!");
}

void subscribe_mqtt() {
  char topic[32];
  sprintf(topic, "ioh/client/%s/hub", HOSTNAME);
  client.subscribe(topic);
  sprintf(topic, "ioh/client/%s/config", HOSTNAME);
  client.subscribe(topic);
}

void setup() {
  Serial.begin(9600);
  pinMode(LED_PIN, OUTPUT);
  digitalWrite(LED_PIN, LOW);
  
  String(ESP.getChipId(), HEX).toCharArray(HOSTNAME, 8);
  
  connect_wifi();
  connect_mqtt();
  subscribe_mqtt();
  
  current_state = EMPTY;
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.println(topic);
  if (strstr(topic,"/hub") && current_state == EMPTY) {
    current_state = ACKED;
    Serial.println("ACKED");
  }
  if (strstr(topic,"/config")) {
    char to_parse[16];
    strncpy(to_parse, (char*)payload, length);
    to_parse[length] = '\0';
    char *str_end;
    current_config.water = atoi((char*)to_parse);
    current_state = RUNNING;
    Serial.print("UPDATED CONFIG: ");
    Serial.println(current_config.water);
  }
}

void do_empty() {
  if (millis() - current_config.last_tick > 1000 && client.connected()) {
    char topic[32];
    sprintf(topic, "ioh/client/%s/discover", HOSTNAME);
    client.publish(topic, "EMPTY");
    current_config.last_tick = millis();
  }
}

void do_running() {
  if (millis() - current_config.last_tick > current_config.water * 100) {
    digitalWrite(LED_PIN, current_config.pin_is_on ? LOW : HIGH);
    current_config.pin_is_on = !current_config.pin_is_on;
    current_config.last_tick = millis();
  }
}

void loop() {
  STATE prev = current_state;
  if (!client.connected()) {
    if (millis() - LAST_CONNECT > 1000) {
      digitalWrite(LED_PIN, LOW);
      if (client.connect(HOSTNAME, NULL, NULL)) {
        subscribe_mqtt();
      }
    }
  }
  else {
    client.loop();
  }
  switch(current_state) {
    case EMPTY:
      do_empty();
      break;
    case RUNNING:
      do_running();
      break;
  }
  if (prev != current_state){
    Serial.print("New state: ");
    Serial.println(current_state);
  }
}
