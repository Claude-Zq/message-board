package dao

import "message-board/model"

func InsertPost(post model.Post) error {
	_, err := dB.Exec("INSERT INTO post(username,txt,post_time,update_time) values(?,?,?,?);", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}

func SelectPostById(postId int) (model.Post, error) {
	post := model.Post{}

	row := dB.QueryRow("SELECT id,username,comment_num,txt,post_time,update_time FROM post WHERE id = ?", postId)

	if row.Err() != nil {
		return model.Post{}, row.Err()
	}

	err := row.Scan(&post.Id, &post.Username, &post.CommentNum, &post.Txt, &post.PostTime, &post.UpdateTime)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}
