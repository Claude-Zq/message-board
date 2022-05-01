package dao_test

import (
	"message-board/dao"
	"message-board/model"
	"testing"
	"time"
)

func TestInsertComment(t *testing.T) {
	dao.InitDB()
	comment := model.Comment{
		PostId:      2,
		Txt:         "这是测试代码插入的评论",
		Username:    "admin",
		CommentTime: time.Now(),
	}
	err := dao.InsertComment(comment)
	if err != nil {
		t.Log(err)
	}
}
