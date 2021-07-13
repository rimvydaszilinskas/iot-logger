package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/iot-logger/models"
)

func (app *App) StoreDeviceState() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		device := ctx.MustGet(ContextDeviceKey).(*models.Device)
		var system models.FullSystemDetails

		if err := ctx.BindJSON(&system); err != nil {
			log.Printf("Bad request data, error: %s", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"detail": "Bad request data",
			})
			return
		}

		err := app.redis.StoreDeviceState(&system, device)

		if err != nil {
			fmt.Println(err)
		}

		ctx.Status(http.StatusNoContent)
	}
}
