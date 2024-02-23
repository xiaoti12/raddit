package redisdb

import (
	"github.com/redis/go-redis/v9"
	"raddit/models"
)

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
