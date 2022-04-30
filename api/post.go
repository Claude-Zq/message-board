package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postDetail(ctx *gin.Context) {
	//postIdString := ctx.Param("post_id")
	//postId, err := strconv.Atoi(postIdString)
	//if err != nil {
	//	fmt.Println("post id string to int err:", err)
	//	tool.RespErrorWithDate(ctx,"post_id格式有误")
	//	return
	//}
	//
	////根据postId拿到post
	//post,err :=
	ctx.String(http.StatusOK, "留言详情")
	return
}

func briefPosts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "留言概略")
	return
}

func addPost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加留言")
	return
}
