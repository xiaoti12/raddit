package redisdb

import (
	"github.com/redis/go-redis/v9"
	"raddit/models"
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

func ChangePostScore(userID, postID string, attitude, score float64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, KeyPostScoreZSet, score, postID)
	if attitude == 0 {
		pipeline.ZRem(ctx, KeyPostVotedZSetPrefix+postID, userID)
	} else {
		pipeline.ZAdd(ctx, KeyPostVotedZSetPrefix+postID, redis.Z{
			Score:  attitude,
			Member: userID,
		})
	}
	_, err := pipeline.Exec(ctx)
	return err
}

func GetOrderedPostIDs(p *models.PostListParams) ([]string, error) {
	// judge order type
	key := KeyPostTimeZSet
	if p.OrderType == models.OrderByScore {
		key = KeyPostScoreZSet
	}
	// confirm query start and stop
	start := (p.Page - 1) * p.Size
	stop := start + p.Size - 1

	return rdb.ZRevRange(ctx, key, int64(start), int64(stop)).Result()
}

func GetPostVoteData(ids []string) ([]int64, error) {
	pipeline := rdb.TxPipeline()
	voteCounts := make([]int64, 0, len(ids))
	for _, id := range ids {
		pipeline.ZCount(ctx, KeyPostVotedZSetPrefix+id, "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, cmd := range cmders {
		count := cmd.(*redis.IntCmd).Val()
		voteCounts = append(voteCounts, count)
	}
	return voteCounts, nil
}
