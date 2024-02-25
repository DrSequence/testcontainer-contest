package model

type PortfolioResult struct {
	Name    string `bson:"name" json:"name"`
	Details string `bson:"details" json:"details"`
}
