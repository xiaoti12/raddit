package service

import (
	"go.uber.org/zap"
	"raddit/dao/mysql"
	"raddit/dao/redisdb"
	"raddit/models"
	"raddit/pkg/snowflake"
	"strconv"
)

func CreatePost(post *models.Post) error {
	post.ID = snowflake.GenID()

	err := mysql.InsertPost(post)
	if err != nil {
		return err
	}

	postID := strconv.Itoa(int(post.ID))
	postTime := float64(post.CreateTime.Unix())
	err = redisdb.CreatePostTime(postID, postTime)
	return err
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

func GetPostList(page, size int) ([]*models.PostDetail, error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("get post list in GetPostDetailList() error", zap.Error(err))
		return nil, err
	}
	postDetails := make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		postDetail := completePostInfo(post)
		postDetails = append(postDetails, postDetail)
	}
	return postDetails, nil
}

func GetOrderedPostList(p *models.PostListParams) ([]*models.PostDetail, error) {
	// get id list from redis
	ids, err := redisdb.GetOrderedPostIDs(p)
	if err != nil {
		zap.L().Error("get post ids from redis in GetPostDetailList() error", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("no post from redis in GetPostDetailList()", zap.Any("params", p))
		return nil, nil
	}
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("get post list from mysql in GetPostDetailList() error", zap.Error(err))
		return nil, err
	}
	postDetails := make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		postDetail := completePostInfo(post)
		postDetails = append(postDetails, postDetail)
	}
	return postDetails, nil
}

func completePostInfo(post *models.Post) *models.PostDetail {
	postDetail := new(models.PostDetail)
	postDetail.Post = post

	authorName, err := mysql.GetUsernameByID(postDetail.AuthorID)
	if err != nil {
		zap.L().Error("get author name in GetPostDetailList() error", zap.Error(err), zap.Int64("author_id", postDetail.AuthorID))
		return nil
	}
	postDetail.AuthorName = authorName

	community, err := mysql.GetCommunityBasic(postDetail.CommunityID)
	if err != nil {
		zap.L().Error("get community in GetPostDetailList() error", zap.Error(err), zap.Int64("community_id", postDetail.CommunityID))
		return nil
	}
	postDetail.CommunityName = community.Name

	return postDetail
}
