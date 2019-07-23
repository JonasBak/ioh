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

type ClientConfig struct {
  Plant string
  Water int
}

type IOHConfig struct {
  c *redis.Client
}

func (conf IOHConfig) GetConfig(p string) *ClientConfig {
  val, err := conf.c.Get(fmt.Sprintf("%s%s", CONFIG_PREFIX, p)).Result()
  if err == redis.Nil {
    return nil
  } else if err != nil {
    panic(err)
  }
  var c ClientConfig
  json.Unmarshal([]byte(val), &c)
  return &c
}

func (conf IOHConfig) SetConfig(p string, config ClientConfig) {
  str, err := json.Marshal(config)
  if err != nil {
    panic(err)
  }
  conf.c.Set(fmt.Sprintf("%s%s", CONFIG_PREFIX, p), str, 0)

  conf.c.SRem(UNCONFIGURED, p)
}

func (conf IOHConfig) AddUnconfigured(p string) {
  err := conf.c.SAdd(UNCONFIGURED, p).Err()
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
