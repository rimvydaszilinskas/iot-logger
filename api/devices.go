package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	apimodels "github.com/rimvydaszilinskas/announcer-backend/api/models"
	"github.com/rimvydaszilinskas/announcer-backend/models"
)

func (api *App) deviceListEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(ContextUserKey).(*models.User)
		devices, found, err := api.db.GetUserDevices(user)

		if !found || err != nil {
			if err != nil {
				log.Printf("error retrieving devices - %s", err)
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error",
			})
			return
		}
		ctx.JSON(http.StatusOK, devices)
	}
}

func (api *App) deviceCreationEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var device apimodels.Device
		user := ctx.MustGet(ContextUserKey).(*models.User)

		if err := ctx.ShouldBind(&device); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validationErrors := device.Validate()

		if len(validationErrors) != 0 {
			ctx.JSON(http.StatusBadRequest, validationErrors)
			return
		}

		newDevice := models.Device{
			VerboseName: device.VerboseName,
			Owner:       *user,
			OwnerID:     int(user.ID),
		}
		newDevice.GenerateToken()

		if err := api.db.DB.Create(&newDevice).Error; err != nil {
			log.Printf("error creating new device - %s", err)
			return
		}

		ctx.JSON(http.StatusCreated, newDevice)
	}
}
