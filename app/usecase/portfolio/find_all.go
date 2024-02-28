package portfolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	pm "testcontainer-contest/domain"
)

func (m *MongoPortfolioService) FindAll(ctx context.Context, page, pageSize int) ([]*pm.Portfolio, error) {
	findOptions := options.Find().SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))

	cursor, err := m.collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var portfolios []*pm.Portfolio
	for cursor.Next(ctx) {
		var portfolio pm.Portfolio
		if err := cursor.Decode(&portfolio); err != nil {
			return nil, err
		}
		portfolios = append(portfolios, &portfolio)
	}

	return portfolios, nil
}
