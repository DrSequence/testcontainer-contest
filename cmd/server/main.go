package main

import (
	"fmt"
	"log"
	"net/http"
	"testcontainer-contest/config"

	pt "testcontainer-contest/app/usecase/portfolio"
	"testcontainer-contest/handler"
)

func main() {
	cfg := config.ReadConfig()
	service, err := pt.NewMongoPortfolioService(cfg)

	if err != nil {
		log.Fatal("Failed to create portfolio service:", err)
	}

	http.HandleFunc("/api/v1/portfolio", handler.HandleGetPortfolio(service))
	http.HandleFunc("/api/v1/save-portfolio", handler.HandleSavePortfolio(service))
	http.HandleFunc("/api/v1/portfolios", handler.HandleGetPortfolios(service))
	fmt.Println("Server is running on http://localhost:" + cfg.Server.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, nil))
}
