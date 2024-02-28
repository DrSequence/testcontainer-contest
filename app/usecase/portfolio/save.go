package portfolio

import (
	"context"

	"testcontainer-contest/domain"
)

func (m *MongoPortfolioService) Save(ctx context.Context, portfolio *domain.Portfolio) error {
	_, err := m.collection.InsertOne(ctx, portfolio)
	return err
}
