package portfolio

import (
	"testcontainer-contest/config"

	pm "testcontainer-contest/repository/portfolio/mongo"
)

func NewMongoPortfolioService(cfg config.Config) (*MongoPortfolioService, error) {
	col, err := pm.NewMongoDb(cfg.Database.Address, cfg.Database.Database, cfg.Database.Collection)
	if err != nil {
		return nil, err
	}

	return &MongoPortfolioService{collection: col}, nil
}
