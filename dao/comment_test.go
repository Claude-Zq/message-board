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

func TestSelectCommentByPostId(t *testing.T) {
	dao.InitDB()
	comments, err := dao.SelectCommentByPostId(2)
	if err != nil {
		t.Log(err)
		return
	}
	for _, comment := range comments {
		t.Log(comment)
	}

}

func TestDeleteCommentById(t *testing.T) {
	dao.InitDB()
	err := dao.DeleteCommentById(10)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateCommentByCommentId(t *testing.T) {
	dao.InitDB()
	err := dao.UpdateCommentByCommentId(5, "新评论")
	if err != nil {
		t.Log(err)
	}
}
