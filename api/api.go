package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/db"
)

type App struct {
	db *db.Connection
}

func GetApplication(connection *db.Connection, api *gin.RouterGroup) *App {

	app := App{
		db: connection,
	}

	api.Use(CORSMiddleware())

	api.GET("/", app.pingEndpoint())
	api.POST("/register", app.registrationEndpoint())
	api.POST("/login", app.authenticationEndpoint())

	api.Use(app.tokenMiddleware())

	api.GET("/devices", app.deviceListEndpoint())
	api.POST("/devices", app.deviceCreationEndpoint())
	api.GET("/devices/:id/entries/latest", app.latestDeviceEntry())

	return &app
}

func (app *App) InitIot(iot *gin.RouterGroup) {
	iot.Use(app.iotTokenMiddleware())
	iot.GET("/", app.deviceEntryList())
	iot.POST("/", app.addDeviceEntryEndpoint())
}
