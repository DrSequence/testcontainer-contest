package portfolio

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPortfolioService struct {
	collection *mongo.Collection
}
