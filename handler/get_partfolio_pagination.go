package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	s "testcontainer-contest/app/usecase/portfolio"
)

const (
	pageQuery = "page"
	sizeQuery = "pageSize"
)

func HandleGetPortfolios(service s.PortfolioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get(pageQuery)
		pageSizeStr := r.URL.Query().Get(sizeQuery)

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

		portfolios, err := service.FindAll(r.Context(), page, pageSize)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching portfolios: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(portfolios)
	}
}
