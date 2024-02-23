package redisdb

import (
	"github.com/redis/go-redis/v9"
)

func CreatePostTime(id string, time float64) error {
	pipeline := rdb.TxPipeline()
	// add post time (not changed)
	pipeline.ZAdd(ctx, KeyPostTimeZSet, redis.Z{
		Score:  time,
		Member: id,
	})
	// add post time for score
	pipeline.ZAdd(ctx, KeyPostScoreZSet, redis.Z{
		Score:  time,
		Member: id,
	})
	_, err := pipeline.Exec(ctx)
	return err
}

func GetPostTime(id string) float64 {
	return rdb.ZScore(ctx, KeyPostTimeZSet, id).Val()
}

func GetUserAttitude(userID, postID string) float64 {
	return rdb.ZScore(ctx, KeyPostVotedZSetPrefix+postID, userID).Val()
}

func ChangePostScore(userID, postID string, attitude, score float64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, KeyPostScoreZSet, score, postID)
	pipeline.ZAdd(ctx, KeyPostVotedZSetPrefix+postID, redis.Z{
		Score:  attitude,
		Member: userID,
	})
	_, err := pipeline.Exec(ctx)
	return err
}
