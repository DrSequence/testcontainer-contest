package portfolio

import (
	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPortfolioService struct {
	collection *mongo.Collection
	client     *memcache.Client
}
