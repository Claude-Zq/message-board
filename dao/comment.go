package dao

import "message-board/model"

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
