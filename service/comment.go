package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComment(comment model.Comment) error {
	err := dao.InsertComment(comment)
	if err != nil {
		return err
	}
	post, err := GetPostById(comment.PostId)
	if err != nil {
		return err
	}
	err = dao.UpdatePostCommentNum(post.Id, post.CommentNum+1)
	return err
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
