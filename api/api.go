package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/db"
	"github.com/rimvydaszilinskas/announcer-backend/rds"
)

type App struct {
	db    *db.Connection
	redis *rds.RedisClient
}

func GetApplication(db *db.Connection, redis *rds.RedisClient, api *gin.RouterGroup) *App {

	app := App{
		db:    db,
		redis: redis,
	}

	api.Use(CORSMiddleware())

	api.GET("/", app.pingEndpoint())
	api.POST("/register", app.registrationEndpoint())
	api.POST("/login", app.authenticationEndpoint())

	api.Use(app.tokenMiddleware())

	api.GET("/devices", app.deviceListEndpoint())
	api.POST("/devices", app.deviceCreationEndpoint())
	api.GET("/devices/:deviceID", app.getDeviceState())

	return &app
}

func (app *App) InitIot(iot *gin.RouterGroup) {
	iot.Use(app.iotTokenMiddleware())

	iot.POST("/", app.StoreDeviceState())
}
