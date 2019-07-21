package ioh_config

import (
  "fmt"
  "encoding/json"
  "github.com/go-redis/redis"
)

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func GetConfig() IOHConfig {
  return IOHConfig {getClient()}
}

type PlantConfig struct {
  Something int
}

type IOHConfig struct {
  c *redis.Client
}

func (conf IOHConfig) GetConfig(p string) *PlantConfig {
  val, err := conf.c.Get(fmt.Sprintf("%s%s", CONFIG_PREFIX, p)).Bytes()
  if err == redis.Nil {
    return nil
  } else if err != nil {
    panic(err)
  }
  var config PlantConfig
  json.Unmarshal(val, &config)
  return &config
}

func (conf IOHConfig) SetConfig(p string, config PlantConfig) {
  str, _ := json.Marshal(config)
  conf.c.Set(UNCONFIGURED, str, 0)
}

func (conf IOHConfig) AddUnconfigured(p string) {
  _, err := conf.c.SAdd(UNCONFIGURED, p).Result()
  if err != nil {
    panic(err)
  }
}

func (conf IOHConfig) GetUnconfigured() []string {
  values, err := conf.c.SMembers(UNCONFIGURED).Result()
  if err != nil {
    panic(err)
  }
  return values
}
