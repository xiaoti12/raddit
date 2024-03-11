package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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

func GetOrderedPostIDsByCommunity(p *models.PostListParams) ([]string, error) {
	communityID := int64(*p.CommunityID)
	communityKey := KeyCommunitySetPrefix + strconv.Itoa(int(communityID))
	// generate order key of certain community
	var cacheKey, orderKey string
	if p.OrderType == models.OrderByTime {
		orderKey = KeyPostTimeZSet
	} else {
		orderKey = KeyPostScoreZSet
	}
	cacheKey = orderKey + fmt.Sprintf(":%d", communityID)
	// check cache
	pipeline := rdb.TxPipeline()
	if rdb.Exists(ctx, cacheKey).Val() == 0 {
		pipeline.ZInterStore(ctx, cacheKey, &redis.ZStore{
			Keys:      []string{communityKey, orderKey},
			Aggregate: "MAX",
		})
		pipeline.Expire(ctx, cacheKey, 3600*time.Second)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			zap.L().Error("create zset of certain community error", zap.Error(err), zap.String("key", cacheKey))
			return nil, err
		}
		zap.L().Debug("create zset of certain community", zap.String("key", cacheKey))
	}
	return getPostIDs(cacheKey, p.Page, p.Size)
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

func InsertPost(p *models.Post) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	rdb.LPush(ctx, KeyPostList, data)
}

func GetPostList(page, size int) ([]*models.Post, error) {
	start := int64((page - 1) * size)
	stop := start + int64(size) - 1
	postsData, err := rdb.LRange(ctx, KeyPostList, start, stop).Result()
	if err != nil {
		return nil, err
	}
	posts := make([]*models.Post, len(postsData))
	for i, data := range postsData {
		var post models.Post
		_ = json.Unmarshal([]byte(data), &post)
		posts[i] = &post
	}
	return posts, nil
}
