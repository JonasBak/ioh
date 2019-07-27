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

func listClients(db *sql.DB, q string) []string {
  rows, err := db.Query(q)
  if err != nil {
    panic(err)
  }
  ids := []string{}
  for rows.Next() {
    var (
      id string
    )
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
    ids = append(ids, id)
	}
  return ids
}
