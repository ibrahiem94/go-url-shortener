package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 6 * time.Hour

func Init() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:7777",
		Password: "",
		DB:       0,
	})

	ping, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\n Redis started successfully: ping message = {%s}", ping)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrl(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func RetrieveUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
