package ioh_config

import (
  "database/sql"
  _ "github.com/lib/pq"
)

func GetConfig() IOHConfig {
  connStr := "user=hub dbname=ioh sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
  return IOHConfig {db}
}

type ClientConfig struct {
  Plant string
  Water int
}

type IOHConfig struct {
  db *sql.DB
}

func (conf IOHConfig) GetConfig(p string) *ClientConfig {
  q := "SELECT plant, water FROM configs WHERE clientid = $1"

  var (
    plant string
    water int
  )

  // TODO p to int, or change id type in db.sql
  err := conf.db.QueryRow(q, p).Scan(&plant, &water)
  if err == sql.ErrNoRows {
    return nil
  } else if err != nil {
    panic(err)
  }
  return &ClientConfig{plant, water}
}

func (conf IOHConfig) updateConfig(p string, config ClientConfig) {
  // TODO handle both create and update
  // or have a function for each
}

func (conf IOHConfig) createConfig(p string, config ClientConfig) {
}


func (conf IOHConfig) SetConfig(p string, config ClientConfig) {
  existing := conf.GetConfig(p)
  if existing == nil {
    conf.createConfig(p, config)
  } else {
    conf.updateConfig(p, config)
  }
}

func (conf IOHConfig) AddClient(p string) {
  q := `INSERT INTO client (id)
VALUES ($1)`

  _, err := conf.db.Exec(q, p)
  if err != nil {
    panic(err)
  }
}

func (conf IOHConfig) GetUnconfigured() []string {
  // TODO list all
  return nil
}
