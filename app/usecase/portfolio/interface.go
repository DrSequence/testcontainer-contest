package portfolio

import (
	"context"
	"testcontainer-contest/domain"
)

type PortfolioService interface {
	FindByID(ctx context.Context, ID string) (*domain.Portfolio, error)
	Save(ctx context.Context, portfolio *domain.Portfolio) (string, error)
	FindAll(ctx context.Context, page, pageSize int) ([]*domain.Portfolio, error)
}
