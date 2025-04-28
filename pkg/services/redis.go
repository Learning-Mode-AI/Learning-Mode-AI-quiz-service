package services

import (
	"Learning-Mode-AI-quiz-service/pkg/config"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	logrus "github.com/sirupsen/logrus"
)

var ctx = context.Background()

// Initialize a new Redis client
var rdb *redis.Client

// Initialize the logger
var logger = logrus.New()

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


	// Log Redis initialization
	logger.WithFields(logrus.Fields{
		"redis_host": config.RedisHost,
	}).Info("Initialized Redis client successfully")

  err := rdb.Ping(ctx).Err()

	if err != nil {
		panic(err)
	}
}

// StoreQuizInRedis stores a quiz in Redis with a specified TTL
func StoreQuizInRedis(userID, videoID string, quiz *AIResponse) error {
    key := fmt.Sprintf("quiz:%s:%s", userID, videoID) // Use "quiz:<user_id>:<video_id>" as the key
    data, err := json.Marshal(quiz)
    if err != nil {
        logger.WithFields(logrus.Fields{
            "error": err.Error(),
            "user_id": userID,
            "video_id": videoID,
        }).Error("Failed to marshal quiz data")
        return fmt.Errorf("failed to marshal quiz data: %w", err)
    }

    err = rdb.Set(ctx, key, data, 168*time.Hour).Err() // Store with a 1-week TTL
    if err != nil {
        logger.WithFields(logrus.Fields{
            "error": err.Error(),
            "user_id": userID,
            "video_id": videoID,
            "redis_key": key,
        }).Error("Failed to store quiz in Redis")
        return fmt.Errorf("failed to store quiz in Redis: %w", err)
    }

    logger.WithFields(logrus.Fields{
        "user_id": userID,
        "video_id": videoID,
        "redis_key": key,
    }).Info("Quiz stored successfully in Redis")

    return nil
}

// GetQuizFromRedis fetches a quiz from Redis
func GetQuizFromRedis(userID, videoID string) (*AIResponse, error) {
    key := fmt.Sprintf("quiz:%s:%s", userID, videoID) // Use "quiz:<user_id>:<video_id>" as the key
    val, err := rdb.Get(ctx, key).Result()

    if err == redis.Nil {
        // Key does not exist
        logger.WithFields(logrus.Fields{
            "user_id": userID,
            "video_id": videoID,
            "redis_key": key,
        }).Info("Quiz not found in Redis")
        return nil, nil
    } else if err != nil {
        logger.WithFields(logrus.Fields{
            "error": err.Error(),
            "user_id": userID,
            "video_id": videoID,
            "redis_key": key,
        }).Error("Failed to retrieve quiz from Redis")
        return nil, fmt.Errorf("failed to retrieve quiz from Redis: %w", err)
    }

    var quiz AIResponse
    err = json.Unmarshal([]byte(val), &quiz)
    if err != nil {
        logger.WithFields(logrus.Fields{
            "error": err.Error(),
            "user_id": userID,
            "video_id": videoID,
            "redis_key": key,
        }).Error("Failed to unmarshal quiz data from Redis")
        return nil, fmt.Errorf("failed to unmarshal quiz data from Redis: %w", err)
    }

    logger.WithFields(logrus.Fields{
        "user_id": userID,
        "video_id": videoID,
        "redis_key": key,
    }).Info("Quiz retrieved successfully from Redis")

    return &quiz, nil
}
