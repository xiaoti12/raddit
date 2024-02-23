package models

import "time"

type Post struct {
	ID          int64     `json:"post_id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type PostDetail struct {
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
	Votes         int64  `json:"vote_num"`
	*Post
	//*CommunityBasic
	// Post 和 CommunityBasic 字段存在相同tag内容，冲突会导致该字段确实
}
