package sendreview

import (
	"encoding/json"
	"net/http"

	"server/internal/storage"
)

type ReviewsSender interface {
	GetReviews() ([]storage.Review, error)
}

func SendReviewsHandler(db ReviewsSender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "use method Get", http.StatusMethodNotAllowed)
			return
		}

		reviews, err := db.GetReviews()
		if err != nil {
			http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)
			return
		}
		//Устаналвиваем заголовок
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(reviews); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}
