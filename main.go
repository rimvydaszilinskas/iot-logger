package main

import (
	"log"

	"github.com/rimvydaszilinskas/iot-logger/db"
	"github.com/rimvydaszilinskas/iot-logger/rds"
	"github.com/rimvydaszilinskas/iot-logger/web"
)

func main() {
	conn, err := db.NewConnection()

	if err != nil {
		log.Fatalf("error opening database connection - %s", err)
	}

	if err := conn.MigrateAll(); err != nil {
		log.Fatalf("error migrating models - %s", err)
	}

	redis, err := rds.GetRedisClient()

	if err != nil {
		log.Fatalf("error connecting to redis - %s", err)
	}

	app := web.GetApplication(conn, redis)

	app.Run()
}
