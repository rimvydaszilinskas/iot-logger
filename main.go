package main

import (
	"github.com/rimvydaszilinskas/announcer-backend/db"
	"github.com/rimvydaszilinskas/announcer-backend/web"
)

func main() {
	conn, err := db.NewConnection()

	if err != nil {
		panic(err)
	}

	if err := conn.MigrateAll(); err != nil {
		panic(err)
	}

	app := web.GetApplication(conn)

	app.Run()
}
