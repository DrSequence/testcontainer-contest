package main

import (
	"fmt"
	"log"
	"net/http"

	pt "testcontainer-contest/app/usecase/portfolio"
	"testcontainer-contest/handler"
)

func main() {
	uri := "mongodb://localhost:27017"
	database := "contest"
	collection := "portfolio"

	service, err := pt.NewMongoPortfolioService(uri, database, collection)
	if err != nil {
		log.Fatal("Failed to create portfolio service:", err)
	}

	http.HandleFunc("/api/v1/portfolio", handler.HandleGetPortfolio(service))
	http.HandleFunc("/api/v1/save-portfolio", handler.HandleSavePortfolio(service))
	http.HandleFunc("/api/v1/portfolios", handler.HandleGetPortfolios(service))
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
