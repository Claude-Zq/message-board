package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)

	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("postId string to int err", err)
		tool.RespErrorWithData(ctx, "文章id有误")
		return
	}

	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		Username:    username,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func deleteComment(ctx *gin.Context) {
	commentIdString := ctx.Param("comment_id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		fmt.Println("comment_id to string err:", err)
		tool.RespErrorWithData(ctx, "comment_id格式有误")
		return
	}
	err = service.DeleteComment(commentId)
	if err != nil {
		if err == dao.ErrCommentNotExist {
			tool.RespErrorWithData(ctx, "评论不存在")
			return
		}
		fmt.Println("delete comment by comment_id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, "删除成功")
}
