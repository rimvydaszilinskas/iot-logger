package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/db"
)

type API struct {
	db *db.Connection
}

func GetApplication(connection *db.Connection) *gin.Engine {
	r := gin.Default()

	api := API{
		db: connection,
	}

	r.Use(CORSMiddleware())

	r.GET("/", api.pingEndpoint())
	r.POST("/register", api.registrationEndpoint())
	r.POST("/login", api.authenticationEndpoint())

	iot := r.Group("/iot")
	iot.Use(api.iotTokenMiddleware())
	iot.GET("/", api.deviceEntryList())
	iot.POST("/", api.addDeviceEntryEndpoint())

	r.Use(api.tokenMiddleware())

	r.GET("/devices", api.deviceListEndpoint())
	r.POST("/devices", api.deviceCreationEndpoint())
	r.GET("/devices/:id/entries/latest", api.latestDeviceEntry())

	return r
}
