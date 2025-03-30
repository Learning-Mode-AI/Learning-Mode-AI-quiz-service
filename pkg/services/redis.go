package services

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"
    "Learning-Mode-AI-quiz-service/pkg/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Initialize a new Redis client
var rdb *redis.Client

func InitRedis(redisHost string, redisPassword string, redisDB int) {
	var tlsConfig *tls.Config
	if config.TLSEnabled {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	} else {
		tlsConfig = nil
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:      config.RedisHost, // Replace with Redis server address
		TLSConfig: tlsConfig,
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
}

// StoreQuizInRedis stores a quiz in Redis with a specified TTL
func StoreQuizInRedis(videoID string, quiz *AIResponse) error {
    key := fmt.Sprintf("quiz:%s", videoID) // Use "quiz:<video_id>" as the key
    data, err := json.Marshal(quiz)
    if err != nil {
        return fmt.Errorf("failed to marshal quiz data: %w", err)
    }

    err = rdb.Set(ctx, key, data, 168*time.Hour).Err() // Store with a 1-week TTL
    if err != nil {
        return fmt.Errorf("failed to store quiz in Redis: %w", err)
    }

    return nil
}


// GetQuizFromRedis fetches a quiz from Redis
func GetQuizFromRedis(videoID string) (*AIResponse, error) {
    key := fmt.Sprintf("quiz:%s", videoID) // Use "quiz:<video_id>" as the key
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        // Key does not exist
        return nil, nil
    } else if err != nil {
        return nil, fmt.Errorf("failed to retrieve quiz from Redis: %w", err)
    }

    var quiz AIResponse
    err = json.Unmarshal([]byte(val), &quiz)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal quiz data from Redis: %w", err)
    }

    return &quiz, nil
}

