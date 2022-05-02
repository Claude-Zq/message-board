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
	//得到评论信息
	comment, err := dao.GetCommentByCommentId(id)
	if err != nil {
		return err
	}
	//删除评论
	err = dao.DeleteCommentByCommentId(id)
	if err != nil {
		return err
	}

	//更新对应留言的评论数
	post, err := dao.SelectPostById(comment.PostId)
	if err != nil {
		return err
	}
	err = dao.UpdatePostCommentNum(comment.PostId, post.CommentNum-1)
	return err
}

func UpdateComment(commentId int, newTxt string) error {
	return dao.UpdateCommentByCommentId(commentId, newTxt)
}
