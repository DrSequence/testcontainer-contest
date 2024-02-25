package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testcontainer-contest/entity"
	"time"

	s "testcontainer-contest/app/usecase/portfolio"
	mapper "testcontainer-contest/pkg"
)

func HandleGetPortfolio(service s.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("ID")
		if name == "" {
			http.Error(w, "Missing ID parameter", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		portfolio, err := service.FindByID(ctx, name)

		res := mapper.MapPortfolioToResult(portfolio)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error finding portfolio: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	}
}

func HandleSavePortfolio(service s.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var portfolio entity.Portfolio
			if err := json.NewDecoder(r.Body).Decode(&portfolio); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := service.Save(ctx, &portfolio); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		}
	}
}

func HandleGetPortfolios(service s.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлечение параметров запроса
		pageStr := r.URL.Query().Get("page")
		pageSizeStr := r.URL.Query().Get("pageSize")

		// Преобразование параметров в числа
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
			return
		}

		// Вызов метода FindAll с использованием параметров пагинации
		portfolios, err := service.FindAll(r.Context(), page, pageSize)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching portfolios: %v", err), http.StatusInternalServerError)
			return
		}

		// Сериализация результатов в JSON и отправка клиенту
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(portfolios)
	}
}
