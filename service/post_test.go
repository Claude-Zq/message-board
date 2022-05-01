package service_test

import (
	"message-board/dao"
	"message-board/model"
	"message-board/service"
	"testing"
	"time"
)

func TestAddPost(t *testing.T) {

	dao.InitDB()
	p := model.Post{
		CommentNum: 0,
		Txt:        "这是李四的留言",
		Username:   "李四",
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.AddPost(p)
	if err != nil {
		t.Log(err)
		return
	}
}

func TestGetPosts(t *testing.T) {
	dao.InitDB()
	posts, err := service.GetPosts()
	if err != nil {
		t.Log(err)
	}
	for index, post := range posts {
		t.Log(index, post)
	}
}

func TestGetPostById(t *testing.T) {

	dao.InitDB()
	if post, err := service.GetPostById(3); err != nil {
		t.Log(err)
	} else {
		t.Log(post)
	}

	if post, err := service.GetPostById(-1); err != nil {
		t.Log(err)
	} else {
		t.Log(post)
	}
}
