package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/api"
	"github.com/rimvydaszilinskas/announcer-backend/db"
	"github.com/rimvydaszilinskas/announcer-backend/rds"
)

type App struct {
	DB    *db.Connection
	Redis *rds.RedisClient
}

func GetApplication(connection *db.Connection, redis *rds.RedisClient) *gin.Engine {
	router := gin.Default()

	app := App{
		DB:    connection,
		Redis: redis,
	}

	apiGroup := router.Group("/api")
	iotGroup := router.Group("/iot")
	api := api.GetApplication(connection, redis, apiGroup)
	api.InitIot(iotGroup)

	router.LoadHTMLGlob("templates/*")

	router.GET("/", app.IndexPage())
	// router.GET("/login", app.LoginPage())
	// router.POST("/login", app.AttemptLogin())

	return router
}
