package portfolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	pm "testcontainer-contest/entity"
)

func (m *MongoPortfolioService) FindByID(ctx context.Context, ID string) (*pm.Portfolio, error) {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	var portfolio pm.Portfolio
	filter := bson.M{"_id": objectID}
	err = m.collection.FindOne(ctx, filter).Decode(&portfolio)
	if err != nil {
		return nil, err
	}

	return &portfolio, nil
}
