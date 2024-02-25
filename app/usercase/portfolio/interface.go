package portfolio

import (
	"context"
	"testcontainer-contest/entity"
)

type Service interface {
	FindByID(ctx context.Context, ID string) (*entity.Portfolio, error)
	Save(ctx context.Context, portfolio *entity.Portfolio) error
	FindAll(ctx context.Context, page, pageSize int) ([]*entity.Portfolio, error)
}
