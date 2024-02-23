package redisdb

import "raddit/models"

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
