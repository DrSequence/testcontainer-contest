package pkg

import (
	e "testcontainer-contest/entity"
	h "testcontainer-contest/model"
)

func MapPortfolioToResult(portfolio *e.Portfolio) *h.PortfolioResult {
	return &h.PortfolioResult{
		Name:    portfolio.Name,
		Details: portfolio.Details,
	}
}
