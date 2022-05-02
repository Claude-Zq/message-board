package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComment(comment model.Comment) error {
	return dao.InsertComment(comment)
}

func GetPostComments(postId int) ([]model.Comment, error) {
	return dao.SelectCommentByPostId(postId)
}

func DeleteComment(id int) error {
	return dao.DeleteCommentByCommentId(id)
}

func UpdateComment(commentId int, newTxt string) error {
	return dao.UpdateCommentByCommentId(commentId, newTxt)
}
