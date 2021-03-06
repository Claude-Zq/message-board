package dao_test

import (
	"fmt"
	"message-board/dao"
	"message-board/model"
	"testing"
	"time"
)

func TestSelectPostById(t *testing.T) {
	dao.InitDB()
	p := model.Post{}
	p, _ = dao.SelectPostById(1)
	t.Log(p)
}

func TestInsertPost(t *testing.T) {
	dao.InitDB()
	p := model.Post{
		CommentNum: 0,
		Txt:        "这是张三的留言",
		Username:   "zhangsan",
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.InsertPost(p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSelectPosts(t *testing.T) {
	dao.InitDB()
	posts, err := dao.SelectPosts()
	if err != nil {
		t.Log(err)
	}
	for _, post := range posts {
		fmt.Println(post)
	}
}

func TestDeletePost(t *testing.T) {
	dao.InitDB()
	err := dao.DeletePost(7)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdatePostTxt(t *testing.T) {
	dao.InitDB()
	err := dao.UpdatePostTxt(2, "测试dao层")
	if err != nil {
		t.Log(err)
	}
}

func TestUpdatePostCommentNum(t *testing.T) {
	dao.InitDB()
	err := dao.UpdatePostCommentNum(1, 2)
	if err != nil {
		t.Log(err)
	}
}

func TestAddLike(t *testing.T) {
	dao.InitDB()
	for i := 0; i < 100; i++ {
		err := dao.AddLike(i)
		if err != nil {
			t.Log(err)
		}
	}
}

func TestDeleteLike(t *testing.T) {
	dao.InitDB()
	for i := 0; i < 100; i++ {
		err := dao.DeleteLike(i)
		if err != nil {
			t.Log(err)
		}
	}
}
