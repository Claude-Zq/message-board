package dao

import (
	"message-board/model"
)

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

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := dB.Query("SELECT id,username,txt,post_time,update_time,comment_num FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil

}
