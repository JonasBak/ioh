package ioh_config

import (
  "database/sql"
)

func exec(db *sql.DB, q string, args ...interface{}) {
  _, err := db.Exec(q, args...)
  if err != nil {
    panic(err)
  }
}

func listClients(db *sql.DB, q string) []Client {
  rows, err := db.Query(q)
  if err != nil {
    panic(err)
  }
  clients := []Client{}
  for rows.Next() {
    var client Client
    if err := rows.Scan(&client.Id, &client.Active); err != nil {
      panic(err)
    }
    clients = append(clients, client)
  }
  return clients
}
