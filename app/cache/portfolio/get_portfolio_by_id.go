package portfolio

import (
	"context"
	"encoding/json"
	"testcontainer-contest/pkg/hash"

	"testcontainer-contest/domain"
)

func (r *RedisCacheService) FindByID(ctx context.Context, ID string) (*domain.Portfolio, error) {
	key := hash.HashWithByteShift(ID)
	result, err := r.client.Get(ctx, key).Result()

	if len(result) == 0 {
		return r.delegate.FindByID(ctx, ID)
	}

	var portfolio domain.Portfolio
	err = json.Unmarshal([]byte(result), &portfolio)
	if err != nil {
		return nil, err
	}

	return &portfolio, nil
}
