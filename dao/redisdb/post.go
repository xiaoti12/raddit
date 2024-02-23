package redisdb

import (
	"github.com/redis/go-redis/v9"
	"raddit/models"
	"strconv"
	"time"
)

func getPostIDs(key string, page, size int) ([]string, error) {
	start := (page - 1) * size
	stop := start + size - 1
	return rdb.ZRevRange(ctx, key, int64(start), int64(stop)).Result()
}

func CreatePostData(postID, communityID string, time float64) error {
	pipeline := rdb.TxPipeline()
	// add post time (not changed)
	pipeline.ZAdd(ctx, KeyPostTimeZSet, redis.Z{
		Score:  time,
		Member: postID,
	})
	// add post time for score
	pipeline.ZAdd(ctx, KeyPostScoreZSet, redis.Z{
		Score:  time,
		Member: postID,
	})
	// add post in community
	pipeline.SAdd(ctx, KeyCommunitySetPrefix+communityID, postID)
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
	return getPostIDs(key, p.Page, p.Size)
}

func GetOrderedPostIDsByCommunity(p *models.PostListParams, communityID int64) ([]string, error) {
	communityKey := KeyCommunitySetPrefix + strconv.Itoa(int(communityID))
	// generate order key of certain community
	var orderKey string
	if p.OrderType == models.OrderByTime {
		orderKey = KeyPostTimeZSet + strconv.Itoa(int(communityID))
	} else {
		orderKey = KeyPostScoreZSet + strconv.Itoa(int(communityID))
	}
	// check cache
	pipeline := rdb.TxPipeline()
	if rdb.Exists(ctx, orderKey).Val() == 0 {
		pipeline.ZInterStore(ctx, orderKey, &redis.ZStore{
			Keys:      []string{communityKey, orderKey},
			Aggregate: "MAX",
		})
		pipeline.Expire(ctx, orderKey, 3600*time.Second)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return getPostIDs(orderKey, p.Page, p.Size)
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
