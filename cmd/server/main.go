package main

import (
	"fmt"
	"log"
	"net/http"

	"testcontainer-contest/app/cache/portfolio"
	pt "testcontainer-contest/app/usecase/portfolio"
	"testcontainer-contest/config"
	"testcontainer-contest/handler"
)

func main() {
	cfg := config.ReadConfig()

	mongoService, err := pt.NewMongoPortfolioService(cfg)
	if err != nil {
		log.Fatal("Failed to create portfolio service:", err)
	}

	cachedProxy, err := portfolio.NewRedisCacheService(cfg, mongoService)
	if err != nil {
		log.Fatal("Failed to create cached portfolio service:", err)
	}

	http.HandleFunc("/api/v1/portfolio", handler.HandleGetPortfolio(cachedProxy))
	http.HandleFunc("/api/v1/save-portfolio", handler.HandleSavePortfolio(cachedProxy))
	http.HandleFunc("/api/v1/portfolios", handler.HandleGetPortfolios(cachedProxy))
	fmt.Println("Server is running on http://localhost:" + cfg.Server.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, nil))
}
