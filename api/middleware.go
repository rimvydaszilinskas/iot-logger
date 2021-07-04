package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ContextUserKey = "AuthTokenUser"
const ContextDeviceKey = "AuthTokenDevice"

func (api *API) tokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authenticationHeader := ctx.Request.Header.Get("Authorization")

		if len(authenticationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"detail": "Token not provided",
			})
			return
		}

		user, found, err := api.db.GetUserByToken(authenticationHeader)

		if !found || err != nil {
			if err != nil {
				log.Printf("error occured retrieving user by token - %s", err)
			}

			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"detail": "token authentication failed",
			})
			return
		}

		ctx.Set(ContextUserKey, user)
	}
}

// CORSMiddleware allows CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (api *API) iotTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authenticationHeader := ctx.Request.Header.Get("Authorization")

		if len(authenticationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"detail": "Token not provided",
			})
			return
		}

		device, found, err := api.db.GetDeviceByToken(authenticationHeader)

		if !found || err != nil {
			if err != nil {
				log.Printf("error occured retrieving device by token - %s", err)
			}

			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"detail": "token authentication failed",
			})
			return
		}

		ctx.Set(ContextDeviceKey, device)
	}
}
