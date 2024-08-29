package redis

import (
	"context"
	"fmt"
	"food-delivery/internal/configs"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var (
	instance *redis.Client
	once     sync.Once
)

// DB return database connection
func DB(cfg *configs.Redis) (*redis.Client, error) {
	var err error
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		})

		err = instance.Ping(context.Background()).Err()
	})

	if err != nil {
		return nil, errors.Wrap(err, "redis.Connect")
	}

	fmt.Println("redis connect", instance.Ping(context.Background()).Err())

	return instance, nil
}
