package portfolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"testcontainer-contest/domain"
)

func (m *MongoPortfolioService) Save(ctx context.Context, portfolio *domain.Portfolio) (string, error) {
	insertResult, err := m.collection.InsertOne(ctx, portfolio)
	if err != nil {
		return "", err
	}

	savedObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("InsertedID is not an ObjectID")
	}

	return savedObjectID.Hex(), err
}
