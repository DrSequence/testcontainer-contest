package portfolio

import (
	"context"
	"encoding/json"
	"testcontainer-contest/pkg/hash"

	"testcontainer-contest/domain"
)

func (r *RedisCacheService) Save(ctx context.Context, portfolio *domain.Portfolio) (string, error) {
	savedObjectID, err := r.delegate.Save(ctx, portfolio)

	jsonPortfolio, err := json.Marshal(portfolio)
	if err != nil {
		return savedObjectID, err
	}

	key := hash.HashWithByteShift(savedObjectID)

	err = r.client.Set(ctx, key, jsonPortfolio, r.expiration).Err()
	if err != nil {
		return savedObjectID, err
	}

	return savedObjectID, nil
}
