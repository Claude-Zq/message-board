package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		ctx.String(http.StatusOK, "请先登陆")
		ctx.Abort()
	}
	ctx.Set("username", username)
	ctx.Next()
}
