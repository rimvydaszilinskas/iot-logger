package main

import (
	"github.com/rimvydaszilinskas/announcer-backend/api"
	"github.com/rimvydaszilinskas/announcer-backend/db"
)

func main() {
	conn, err := db.NewConnection()

	if err != nil {
		panic(err)
	}

	if err := conn.MigrateAll(); err != nil {
		panic(err)
	}

	app := api.GetApplication(conn)

	app.Run()
}
