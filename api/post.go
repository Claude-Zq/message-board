package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"os"
	"path"
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

func like(ctx *gin.Context) {
	postIdString := ctx.Query("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err:", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}
	err = service.Like(postId)
	if err != nil {
		fmt.Println("AddLike to post err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, "点赞成功")

}

func cancelLike(ctx *gin.Context) {
	postIdString := ctx.Query("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err:", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}
	err = service.CancelLike(postId)
	if err != nil {
		fmt.Println("cancelLike to post err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.ResSuccessfulWithData(ctx, "已取消点赞")
}

func uploadPicture(ctx *gin.Context) {
	postIdString := ctx.PostForm("post_id")
	_, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post_id string to int err:", err)
		tool.RespErrorWithData(ctx, "post_id格式有误")
		return
	}

	picture, err := ctx.FormFile("picture")
	if err != nil {
		tool.RespErrorWithData(ctx, "文件上传失败:"+err.Error())
	}
	//获取后缀名
	extName := path.Ext(picture.Filename)
	//判断后缀名是否合法
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".gif":  true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		tool.RespErrorWithData(ctx, "不支持该图片格式,头像上传失败")
		return
	}

	//创建文件保存目录(按留言编号创建)
	dir := "./static/upload/" + postIdString
	//判断文件夹是否已存在
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) /*不存在*/ {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				fmt.Println("创建目录%s失败："+err.Error(), dir)
				return
			}
		} else /*不确定是否存在*/ {
			log.Printf("不能确定目录%s是否存在\n", dir)
			return
		}
	}
	dst := dir + "/" + picture.Filename
	_ = ctx.SaveUploadedFile(picture, dst)
	tool.ResSuccessfulWithData(ctx, "图片上传成功")

}
