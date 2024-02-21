package service

import (
	"go.uber.org/zap"
	"raddit/dao/mysql"
	"raddit/models"
	"raddit/pkg/snowflake"
)

func CreatePost(post *models.Post) error {
	post.ID = snowflake.GenID()
	return mysql.InsertPost(post)
}

func GetPostDetail(id int64) (*models.PostDetail, error) {
	postDetail := new(models.PostDetail)
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("get post detail in GetPostDetail() error", zap.Error(err))
		return nil, err
	}
	postDetail.Post = post

	authorName, err := mysql.GetUsernameByID(postDetail.AuthorID)
	if err != nil {
		zap.L().Error("get author name in GetPostDetail() error", zap.Error(err))
		return nil, err
	}
	postDetail.AuthorName = authorName

	community, err := mysql.GetCommunityBasic(postDetail.CommunityID)
	if err != nil {
		zap.L().Error("get community in GetPostDetail() error", zap.Error(err))
		return nil, err
	}
	postDetail.CommunityName = community.Name

	return postDetail, nil
}
