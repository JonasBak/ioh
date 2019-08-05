#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <WiFiClientSecure.h>

// generated with 'openssl x509 -pubkey -noout -in cert.pem'
static const char pubkey[] PROGMEM = R"KEY(
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArmB24OFMGjGhYiR+goMf
tymcL+n+GNZiSKGmjJHoEOgTMrth3ZaGPSLR7a5u1Lu1qZ2aoM04Oqb+56OpPByw
7ZoJ+/p4/4iliGyPaO8JKfVICievtAA6EwNVjbnfH9LbZBRqzbeaZ70dFrjvFt6/
bQds3XK71jVS27Mqo8d1d5NmoDw6BkabOpgnc/eBPEKsM2YfBv7Ah1oCaBYp8g2i
XMAQgKS1wc4ojOjxnl8bGAgGhmbysJvcjSo+TyaHGcUVONWxzTWu/Qkuay/yo6yf
NpYaTgUjqvxJ1UlwfiLob3n5qd5ox9GdVXTD7MLNRwahjEc/wDPgmpGjzcXm5roo
OwIDAQAB
-----END PUBLIC KEY-----
)KEY";

BearSSL::WiFiClientSecure espClient;
BearSSL::PublicKey key(pubkey);
PubSubClient client(espClient);

const char* WIFI_SSID = "Bakkens nett";
const char* WIFI_PASSWORD = "";
const char* MQTT_BROKER = "mqtt.jbakken.com";
const int   MQTT_PORT = 6543;
const char* MQTT_USER = "";
const char* MQTT_PASSWORD = "";

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

void connect_wifi_initial() {
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

void connect_mqtt_initial() {
  Serial.print("Connecting to MQTT...");
  client.setServer(MQTT_BROKER, MQTT_PORT);
  client.setCallback(callback);
  while (!client.connected()) {
    if (!connect_mqtt()) {
      Serial.print(".");
      delay(500);
    }
  }
  Serial.println("\nConnected!");
}

void set_cert_initial() {
  espClient.setKnownKey(&key);
}

bool connect_mqtt() {
  // TODO set all topics in setup
  char status_topic[32];
  sprintf(status_topic, "ioh/client/%s/status", HOSTNAME);
  if (!client.connect(HOSTNAME, MQTT_USER, MQTT_PASSWORD, status_topic, 0, 1, "OFF")) {
    return false;
  }
  char topic[32];
  sprintf(topic, "ioh/client/%s/hub", HOSTNAME);
  client.subscribe(topic);
  sprintf(topic, "ioh/client/%s/config", HOSTNAME);
  client.subscribe(topic);

  client.publish(status_topic, "ON");
  return true;
}

void setup() {
  Serial.begin(9600);
  pinMode(LED_PIN, OUTPUT);
  digitalWrite(LED_PIN, LOW);

  String(ESP.getChipId(), HEX).toCharArray(HOSTNAME, 8);

  connect_wifi_initial();
  set_cert_initial();
  connect_mqtt_initial();

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
    if (millis() - LAST_CONNECT > 5000) {
      digitalWrite(LED_PIN, LOW);
      connect_mqtt();
      LAST_CONNECT = millis();
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
