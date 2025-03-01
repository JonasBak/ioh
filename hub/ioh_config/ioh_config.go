package ioh_config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func GetConfig() IOHConfig {
	connStr := fmt.Sprintf("user=%s password=%s dbname=ioh host=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return IOHConfig{db}
}

type Client struct {
	Id         string `json:"id"`
	Active     bool   `json:"active"`
	config_ptr *IOHConfig
}

type ClientConfig struct {
	Plant     string `json:"plant"`
	Water     int    `json:"water"`
	Timestamp string `json:"timestamp"`
}

func (conf ClientConfig) ToString() string {
	return fmt.Sprintf("%d", conf.Water)
}

type IOHConfig struct {
	db *sql.DB
}

func (conf IOHConfig) Connected() bool {
	err := conf.db.Ping()
	return err == nil
}

func (conf IOHConfig) GetClient(p string) *Client {
	q := "SELECT id, active FROM clients WHERE id = $1"

	var client Client

	err := conf.db.QueryRow(q, p).Scan(&client.Id, &client.Active)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	client.config_ptr = &conf
	return &client
}

func (conf IOHConfig) GetConfig(p string) *ClientConfig {
	q := "SELECT plant, water, timestamp FROM configs WHERE clientid = $1"

	var (
		plant     string
		water     int
		timestamp string
	)

	err := conf.db.QueryRow(q, p).Scan(&plant, &water, &timestamp)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return &ClientConfig{plant, water, timestamp}
}

func (conf IOHConfig) updateConfig(p string, config ClientConfig) {
	q := `UPDATE configs SET plant = $1, water = $2 WHERE clientid = $3`
	exec(conf.db, q, config.Plant, config.Water, p)
}

func (conf IOHConfig) createConfig(p string, config ClientConfig) {
	q := `INSERT INTO configs (plant, water, clientid) VALUES ($1, $2, $3)`
	exec(conf.db, q, config.Plant, config.Water, p)
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
	if conf.GetClient(p) != nil {
		return
	}
	q := `INSERT INTO clients (id) VALUES ($1)`
	exec(conf.db, q, p)
}

func (conf IOHConfig) SetActive(p string, value bool) {
	conf.AddClient(p)
	q := `UPDATE clients SET active = $1 WHERE id = $2`
	exec(conf.db, q, value, p)
}

func (conf IOHConfig) GetClients() []Client {
	q := "SELECT clients.id, clients.active FROM clients"
	return listClients(conf.db, q, &conf)
}

func (conf IOHConfig) GetUnconfigured() []Client {
	q := "SELECT clients.id, clients.active FROM clients LEFT JOIN configs ON clients.id = configs.clientid WHERE configs.id IS NULL"
	return listClients(conf.db, q, &conf)
}

func (conf IOHConfig) GetConfigured() []Client {
	q := "SELECT clients.id, clients.active FROM clients LEFT JOIN configs ON clients.id = configs.clientid WHERE NOT configs.id IS NULL"
	return listClients(conf.db, q, &conf)
}
