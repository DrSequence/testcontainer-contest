package portfolio

import (
	"context"
	"testcontainer-contest/entity"
)

func (m *MongoPortfolioService) Save(ctx context.Context, portfolio *entity.Portfolio) error {
	_, err := m.collection.InsertOne(ctx, portfolio)
	return err
}
