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
	err := dao.DeleteCommentByCommentId(10)
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteCommentsByPostId(t *testing.T) {
	dao.InitDB()
	err := dao.DeleteCommentsByPostId(2)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateCommentByCommentId(t *testing.T) {
	dao.InitDB()
	err := dao.UpdateCommentByCommentId(4, "99999")
	if err != nil {
		t.Log(err)
	}
}

func TestGetCommentByCommentId(t *testing.T) {
	dao.InitDB()
	comment, err := dao.GetCommentByCommentId(19)
	t.Log(comment)
	if err != nil {
		t.Log(err)
	}

}
