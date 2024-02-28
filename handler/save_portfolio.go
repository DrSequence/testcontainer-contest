package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"testcontainer-contest/domain"
	"time"

	s "testcontainer-contest/app/usecase/portfolio"
)

func HandleSavePortfolio(service s.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var portfolio domain.Portfolio
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
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
