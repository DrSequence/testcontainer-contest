package portfolio

import (
	"context"
	"testcontainer-contest/domain"
)

type Service interface {
	FindByID(ctx context.Context, ID string) (*domain.Portfolio, error)
	Save(ctx context.Context, portfolio *domain.Portfolio) error
	FindAll(ctx context.Context, page, pageSize int) ([]*domain.Portfolio, error)
}
