package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/announcer-backend/api/models"
)

func (*App) IndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "HomePage",
		})
	}
}

func (*App) LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	}
}

func (app *App) AttemptLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.AuthenticationUser
		if err := ctx.ShouldBind(&user); err != nil {
			panic(err)
		}
		errs := user.Validate(app.DB.DB)
		if len(errs) != 0 {
			log.Printf("there are errors %s", errs)
		}
		ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "hello world"})
	}
}
