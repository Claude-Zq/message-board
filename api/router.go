package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"message-board/tool"
)

func InitEngine() {
	engine := gin.Default()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("lengthOk", tool.LengthOk)
		if err != nil {
			fmt.Println("验证器注册失败")
			return
		}
	}

	engine.POST("/register", register)   //注册
	engine.POST("/login", login)         //登陆
	engine.GET("/question", getQuestion) //获取密保问题
	engine.POST("/answer", judgeAnswer)  //判断密保问题

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码

	}
	postGroup := engine.Group("/post")
	{
		postGroup.Use(auth)                       //认证
		postGroup.POST("/", addPost)              //发布新留言
		postGroup.POST("/:post_id", updatePost)   //修改留言
		postGroup.POST("/picture", uploadPicture) //上传图片
		postGroup.DELETE("/:post_id", deletePost) //删除留言

		postGroup.GET("/", briefPosts)         //查看全部留言概略
		postGroup.GET("/:post_id", postDetail) //查看一条留言详细信息和其下属评论

		postGroup.POST("/like", like)         //点赞
		postGroup.DELETE("/like", cancelLike) //取消点赞
	}

	commentGroup := engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/", addComment)                 //发送评论
		commentGroup.DELETE("/:comment_id", deleteComment) //删除评论
		commentGroup.POST("/update", updateComment)        //修改评论
	}
	engine.Run()

}
