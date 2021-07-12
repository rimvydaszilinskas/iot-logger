package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/api"
	"github.com/rimvydaszilinskas/announcer-backend/db"
)

type App struct {
	db *db.Connection
}

func GetApplication(connection *db.Connection) *gin.Engine {
	router := gin.Default()

	app := App{
		db: connection,
	}

	apiGroup := router.Group("/api")
	iotGroup := router.Group("/iot")
	api := api.GetApplication(connection, apiGroup)
	api.InitIot(iotGroup)

	router.LoadHTMLGlob("templates/*")

	router.GET("/", app.IndexPage())
	// router.GET("/login", app.LoginPage())
	// router.POST("/login", app.AttemptLogin())

	return router
}
