package redisdb

func GetUserAttitude(userID, postID string) float64 {
	return rdb.ZScore(ctx, KeyPostVotedZSetPrefix+postID, userID).Val()
}
