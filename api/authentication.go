package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/api/models"
)

func (api *API) authenticationEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authUser models.AuthenticationUser
		if err := ctx.ShouldBindJSON(&authUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, found, err := api.db.GetUserByEmail(authUser.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		if !found {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		if !user.ComparePasswordHash(authUser.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"validation": "bad credentials",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": user.Token,
		})
	}
}
