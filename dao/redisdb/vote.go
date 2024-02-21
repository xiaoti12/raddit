package redisdb

func GetPostTime(id string) (int64, error) {
	return 0, nil
}

func GetUserAttitude(userID, postID string) (float64, error) {
	return 0, nil
}

func IncrPostScore(userID, postID string, attitude float64) error {
	return nil
}
