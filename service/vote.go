package service

import (
	"errors"
	"math"
	"raddit/dao/redisdb"
	"raddit/models"
	"strconv"
	"time"
)

const (
	PostExpireTime = 7 * 24 * 3600 // 7 days
	PostScoreBase  = 432.0
)

var ErrorVoteTimeExpire = errors.New("vote time expire")

func VotePost(p *models.VoteParams) error {
	// prepare for redis argument
	userID := strconv.Itoa(int(p.UserID))
	postID := strconv.Itoa(int(p.PostID))
	curAtti := float64(p.Attitude)
	// check vote time
	postTime := redisdb.GetPostTime(postID)
	if float64(time.Now().Unix())-postTime > PostExpireTime {
		return ErrorVoteTimeExpire
	}
	// get previous attitude
	// -1: dislike 1: like 0: not vote
	// -1/0 changed to 1 coefficient:1
	// 1/0 changed to -1 coefficient:-1
	// 1/-1 changed to 0 coefficient:-1/1
	preAtti := redisdb.GetUserAttitude(userID, postID)
	coeff := -1.0
	if curAtti > preAtti {
		coeff = 1.0
	}
	diff := math.Abs(curAtti - preAtti)
	// update post score and user attitude
	err := redisdb.ChangePostScore(userID, postID, curAtti, coeff*diff*PostScoreBase)
	if err != nil {
		return err
	}
	return nil
}
