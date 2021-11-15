package redis

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(redisUrl string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
