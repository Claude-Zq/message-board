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
