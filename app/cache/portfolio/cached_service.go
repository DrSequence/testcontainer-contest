package portfolio

import (
	"github.com/go-redis/redis/v8"
	"testcontainer-contest/app/usecase/portfolio"
	"time"
)

type RedisCacheService struct {
	delegate portfolio.PortfolioService

	client     *redis.Client
	expiration time.Duration
}
