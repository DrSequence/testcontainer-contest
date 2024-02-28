package portfolio

import (
	"context"

	"testcontainer-contest/domain"
)

func (r *RedisCacheService) FindAll(ctx context.Context, page, pageSize int) ([]*domain.Portfolio, error) {
	return r.delegate.FindAll(ctx, page, pageSize)
}
