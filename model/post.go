package model

type Post struct {
	Id         int    `json:"id"`
	CommentNum int    `json:"comment_num"`
	Txt        string `json:"txt"`
	Username   string `json:"username"`
	PostTime   string `json:"post_time"`
	UpdateTime string `json:"update_time"`
}
