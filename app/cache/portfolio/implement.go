package portfolio

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testcontainer-contest/app/usecase/portfolio"
	"testcontainer-contest/config"
)

func NewRedisCacheService(cfg config.Config, service *portfolio.MongoPortfolioService) (*RedisCacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Address,
		DB:       0,
		Password: cfg.Cache.Pass,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisCacheService{
		delegate:   service,
		client:     client,
		expiration: cfg.Cache.Exp,
	}, nil
}
