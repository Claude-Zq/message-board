package dao

import (
	"errors"
	"message-board/model"
)

var ErrCommentNotExist = errors.New("评论不存在")

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(post_id,username,comment_time,txt) VALUES(?,?,?,?)", comment.PostId, comment.Username, comment.CommentTime, comment.Txt)
	if err != nil {
		return err
	}
	return nil
}

func SelectCommentByPostId(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := dB.Query("SELECT id,post_id,username,txt,comment_time FROM comment WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Username, &comment.Txt, &comment.CommentTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func DeleteCommentByCommentId(commentId int) error {
	ret, err := dB.Exec("DELETE FROM comment WHERE id = ?", commentId)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrCommentNotExist
	}
	return nil
}

func DeleteCommentsByPostId(PostId int) error {
	_, err := dB.Exec("DELETE FROM comment WHERE post_id = ?", PostId)
	return err
}

func UpdateCommentByCommentId(commentId int, newTxt string) error {
	_, err := dB.Exec("UPDATE comment SET txt = ? WHERE id = ?", newTxt, commentId)
	return err
}
