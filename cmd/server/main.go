package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: nil, // default
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	go func() {
		fmt.Println("Server is running on http://localhost:" + cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	<-shutdownChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
