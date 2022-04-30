package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"message-board/tool"
	"net/http"
)

func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")

	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	//检验旧密码是否正确
	flag, err := service.IsPasswordCorrect(username, oldPassword)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "旧密码错误")
		return
	}
	//修改新密码
	err = service.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password err:", err)
		return
	}
	tool.RespSuccessful(ctx)
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}
	ctx.SetCookie("username", username, 700, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

func register(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户注册")
}
