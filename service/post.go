package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost(post model.Post) error {
	err := dao.InsertPost(post)
	return err
}

func GetPosts() ([]model.Post, error) {
	return dao.SelectPosts()
}

func GetPostById(postId int) (model.Post, error) {
	return dao.SelectPostById(postId)
}

func DeletePost(postId int) error {
	err := dao.DeletePost(postId)
	if err != nil {
		return err
	}
	err = dao.DeleteCommentsByPostId(postId)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(postId int, newTxt string) error {
	return dao.UpdatePostTxt(postId, newTxt)
}

func Like(postId int) error {
	return dao.AddLike(postId)
}

func CancelLike(postId int) error {
	return dao.DeleteLike(postId)
}
