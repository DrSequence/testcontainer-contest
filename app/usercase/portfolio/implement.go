package portfolio

import (
	"go.mongodb.org/mongo-driver/mongo"

	pm "testcontainer-contest/repository/portfolio/mongo"
)

type MongoPortfolioService struct {
	collection *mongo.Collection
}

func NewMongoPortfolioService(uri, database, collection string) (*MongoPortfolioService, error) {
	col, err := pm.NewMongoDb(uri, database, collection)
	if err != nil {
		return nil, err
	}
	return &MongoPortfolioService{collection: col}, nil
}
