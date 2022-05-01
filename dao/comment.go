package dao

import "message-board/model"

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(post_id,username,comment_time,txt) VALUES(?,?,?,?)", comment.PostId, comment.Username, comment.CommentTime, comment.Txt)
	if err != nil {
		return err
	}
	return nil
}
