package sendreview

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"server/internal/logger/sl"
	"server/internal/storage"
)

type ReviewsSender interface {
	GetReviews() ([]storage.Review, error)
}

func SendReviewsHandler(log *slog.Logger, db ReviewsSender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.sendReview.SendReviewsHandler"

		if r.Method != http.MethodGet {
			log.Error("failed method: use GET")

			http.Error(w, "use method Get", http.StatusMethodNotAllowed)

			return
		}

		reviews, err := db.GetReviews()
		if err != nil {
			log.Error("Failed to fetch reviews", sl.Err(err))

			http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)

			return
		}
		//Устаналвиваем заголовок
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(reviews); err != nil {
			log.Error("Failed to encode response", sl.Err(err))

			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}
