package redisdb

import "github.com/redis/go-redis/v9"

func GetPostTime(id string) float64 {
	return rdb.ZScore(ctx, KeyPostTimeZSet, id).Val()
}

func GetUserAttitude(userID, postID string) float64 {
	return rdb.ZScore(ctx, KeyPostVotedZSetPrefix+postID, userID).Val()
}

func ChangePostScore(userID, postID string, attitude, score float64) error {
	_, err := rdb.ZIncrBy(ctx, KeyPostScoreZSet, score, postID).Result()
	if err != nil {
		return err
	}
	_, err = rdb.ZAdd(ctx, KeyPostVotedZSetPrefix+postID, redis.Z{
		Score:  attitude,
		Member: userID,
	}).Result()
	return err
}
