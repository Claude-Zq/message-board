package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/model"
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
	var u model.Login
	if err := ctx.ShouldBind(&u); err != nil {
		tool.RespErrorWithData(ctx, err.Error())
		return
	}
	//username := ctx.PostForm("username")
	//password := ctx.PostForm("password")

	flag, err := service.IsPasswordCorrect(u.Username, u.Password)
	if err != nil {
		fmt.Println("judge password correct err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}
	ctx.SetCookie("username", u.Username, 700, "/", "", false, false)
	tool.RespSuccessful(ctx)
}

func register(ctx *gin.Context) {
	var u model.Login
	if err := ctx.ShouldBind(&u); err != nil {
		tool.RespErrorWithData(ctx, err.Error())
		return
	}

	//username := ctx.PostForm("username")
	//password := ctx.PostForm("password")

	user := model.User{
		Username: u.Username,
		Password: u.Password,
	}
	flag, err := service.IsRepeatUsername(u.Username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}

	if flag {
		tool.RespErrorWithData(ctx, "用户名已存在")
		return
	}
	err = service.Register(user)
	if err != nil {
		fmt.Println("register err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}

func getQuestion(ctx *gin.Context) {
	username := ctx.Query("username")
	if question, err := service.GetQuestionByName(username); err != nil {
		if err == dao.ErrNoQuestion {
			tool.RespErrorWithData(ctx, "密保问题不存在")
			return
		}
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(ctx, "用户不存在")
			return
		}
		fmt.Println("getQuestion err:", err)
		tool.RespInternalError(ctx)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"question": question,
		})
	}
}

func judgeAnswer(ctx *gin.Context) {
	username := ctx.PostForm("username")
	answer := ctx.PostForm("answer")

	flag, err := service.IsCorrectAnswer(username, answer)

	if err != nil {
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(ctx, "用户不存在")
			return
		}
		if err == dao.ErrNoAnswer {
			tool.RespErrorWithData(ctx, "密保答案不存在")
			return
		}
		fmt.Println("judge answer err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithData(ctx, "答案错误")
	}

}
