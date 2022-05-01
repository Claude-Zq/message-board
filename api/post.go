package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func postDetail(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err:", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}
	//根据postId拿到post
	post, err := service.GetPostById(postId)
	if err != nil {
		fmt.Println("get post by id err", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, post)

}

func briefPosts(ctx *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, posts)
}

func addPost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")

	post := model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)

}
