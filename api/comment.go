package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addComment(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加评论")

}
