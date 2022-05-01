package dao_test

import (
	"message-board/dao"
	"testing"
)

func TestGetAnswerByUsername(t *testing.T) {
	dao.InitDB()
	question0, err0 := dao.GetAnswerByUsername("zhangsan")
	t.Log(question0, err0)
	question1, err1 := dao.GetAnswerByUsername("claude")
	t.Log(question1, err1)
	question2, err2 := dao.GetAnswerByUsername("ahhahha")
	t.Log(question2, err2)
}

func TestGetQuestionByUsername(t *testing.T) {
	dao.InitDB()
	question0, err0 := dao.GetQuestionByUsername("zhangsan")
	t.Log(question0, err0)
	question1, err1 := dao.GetQuestionByUsername("claude")
	t.Log(question1, err1)
	question2, err2 := dao.GetQuestionByUsername("ahhahha")
	t.Log(question2, err2)

}

func TestSelectUserByUsername(t *testing.T) {
	dao.InitDB()
	u, err := dao.SelectUserByUsername("zhangsan")
	t.Log(u, err)
}

func TestUpdatePassword(t *testing.T) {
	dao.InitDB()
	err := dao.UpdatePassword("zhangsan", "321")
	if err != nil {
		t.Log(err)
		return
	}
}
