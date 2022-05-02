package api

import (
	"database/sql"
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
		if err == sql.ErrNoRows {
			tool.RespErrorWithData(ctx, "没有对应的留言")
			return
		}
		fmt.Println("get post by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//根据postId拿到对应comments
	comments, err := service.GetPostComments(postId)
	if err != nil {
		fmt.Println("get comment by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	postDetail := model.PostDetail{
		Post:    post,
		Comment: comments,
	}
	tool.ResSuccessfulWithData(ctx, postDetail)

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

func deletePost(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}
	err = service.DeletePost(postId)
	if err != nil {
		fmt.Println("Delete post err by postId err:", err)
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func updatePost(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	newTxt := ctx.PostForm("new_txt")
	fmt.Println(postIdString, newTxt)

	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err:", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}
	err = service.UpdatePost(postId, newTxt)
	if err != nil {
		fmt.Println("Update post err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, "修改成功")

}
