package service

import (
	"go.uber.org/zap"
	"raddit/dao/mysql"
	"raddit/models"
	"raddit/pkg/jwt"
	"raddit/pkg/snowflake"
)

func Register(p *models.RegisterParams) error {
	// check if user already exists
	err := mysql.CheckUserExists(p.Username)
	if err != nil {
		return err
	}
	// generate user id
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// insert user info to database
	err = mysql.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(p *models.LoginParams) (string, error) {
	// check if user exists
	err := mysql.CheckUserNotExists(p.Username)
	if err != nil {
		return "", err
	}
	// check if password is correct
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.CheckUserLogin(user)
	if err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}

func GetPostDetailList(page, size int) ([]*models.PostDetail, error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("get post list in GetPostDetailList() error", zap.Error(err))
		return nil, err
	}
	postDetails := make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		postDetail := new(models.PostDetail)
		postDetail.Post = post

		authorName, err := mysql.GetUsernameByID(postDetail.AuthorID)
		if err != nil {
			zap.L().Error("get author name in GetPostDetailList() error", zap.Error(err), zap.Int64("author_id", postDetail.AuthorID))
			continue
		}
		postDetail.AuthorName = authorName

		community, err := mysql.GetCommunityBasic(postDetail.CommunityID)
		if err != nil {
			zap.L().Error("get community in GetPostDetailList() error", zap.Error(err), zap.Int64("community_id", postDetail.CommunityID))
			continue
		}
		postDetail.CommunityName = community.Name

		postDetails = append(postDetails, postDetail)
	}
	return postDetails, nil
}
