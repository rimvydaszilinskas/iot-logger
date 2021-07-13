package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/iot-logger/api/models"
)

func (api *App) registrationEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.AuthenticationUser
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validationErrors := user.Validate(api.db.DB)

		if len(validationErrors) != 0 {
			ctx.JSON(http.StatusBadRequest, validationErrors)
			return
		}

		dbUser := user.ToUserModel()
		dbUser.GenerateToken()
		api.db.DB.Create(&dbUser)

		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}
