package service

import (
	"raddit/dao/mysql"
	"raddit/models"
	"raddit/pkg/snowflake"
)

func CreatePost(post *models.Post) error {
	post.ID = snowflake.GenID()
	return mysql.InsertPost(post)
}
