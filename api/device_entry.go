package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apimodels "github.com/rimvydaszilinskas/announcer-backend/api/models"
	"github.com/rimvydaszilinskas/announcer-backend/models"
)

func (api *API) deviceEntryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		device := ctx.MustGet(ContextDeviceKey).(*models.Device)
		ctx.JSON(http.StatusOK, device.Entries)
	}
}

func (api *API) addDeviceEntryEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		device := ctx.MustGet(ContextDeviceKey).(*models.Device)
		var entry models.DeviceEntry
		var apiEntry apimodels.DeviceEntry

		if err := ctx.ShouldBindJSON(&apiEntry); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		entry = models.DeviceEntry{
			PublicIP: apiEntry.PublicIP,
			LocalIP:  apiEntry.LocalIP,
		}

		device.Entries = append(device.Entries, entry)

		if err := api.db.DB.Model(device).Association("Entries").Append(&entry); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, device.Entries[len(device.Entries)-1])
	}
}
