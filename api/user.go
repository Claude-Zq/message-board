package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func changePassword(ctx *gin.Context) {
	ctx.String(http.StatusOK, "修改密码")
}

func login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户登陆")
}

func register(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户注册")
}
