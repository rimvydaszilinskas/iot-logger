package db

import (
	"fmt"
	"os"

	"github.com/rimvydaszilinskas/announcer-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection() (*Connection, error) {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	host := os.Getenv("DB_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = "fms"
	}
	username := os.Getenv("DB_USERNAME")
	if len(username) == 0 {
		username = "fms"
	}
	password := os.Getenv("DB_PASSWORD")
	if len(password) == 0 {
		password = "password"
	}
	port := os.Getenv("DB_PORT")
	if len(port) == 0 {
		port = "5433"
	}
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", host, port, dbName, username, password)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}))
	return &Connection{
		DB: db,
	}, err
}

func (c *Connection) MigrateAll() error {
	models := []interface{}{
		&models.Device{},
		&models.User{},
		&models.DeviceEntry{},
	}

	for _, model := range models {
		if err := c.DB.AutoMigrate(model); err != nil {
			return fmt.Errorf("error migrating %s - %s", model, err)
		}
	}

	return nil
}
