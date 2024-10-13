package redisService

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return rdb
}

func StoreQuiz(sessionID string, quiz map[string]interface{}) error {
	rdb := InitRedis()
	return rdb.Set(ctx, sessionID, quiz, 0).Err()
}

func FetchQuiz(sessionID string) (map[string]interface{}, error) {
	rdb := InitRedis()
	quiz := make(map[string]interface{})
	err := rdb.Get(ctx, sessionID).Scan(&quiz)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}
