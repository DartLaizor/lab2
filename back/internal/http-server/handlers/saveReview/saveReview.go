package savereview

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator/v10"

	"server/internal/logger/sl"
	"server/internal/storage"
)

type ReviewsSaver interface {
	SaveRewiev(r storage.Review) error
}

func SaveReviewsHadnler(log *slog.Logger, db ReviewsSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.saveReview.SaveReviewsHadnler"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		if r.Method != http.MethodPost {
			log.Error("failed method: use POST")

			http.Error(w, "Use method POST ", http.StatusMethodNotAllowed)

			return
		}

		var review storage.Review

		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			log.Error("Invalid JSON", sl.Err(err))

			http.Error(w, "Invalid JSON", http.StatusBadRequest)

			return
		}

		if err := validator.New().Struct(review); err != nil {

			log.Error("invalid request", sl.Err(err))

			http.Error(w, "Invalid JSON", http.StatusBadRequest)

			return
		}
		err := db.SaveRewiev(review)
		data, _ := json.MarshalIndent(review, "", "  ")
		log.Info("Received review", slog.String("data", string(data)))

		if err != nil {
			log.Error("couldn't save to the DataBase", sl.Err(err))

			http.Error(w, "couldn't save to the DataBase", http.StatusInternalServerError)

			return
		}

		/* if review.Name == "" || review.Email!="" || review.Phone != "" || len(review.Technologies)!=0 || review.Date == "" || review.Rating < 0 || review.Rating > 9 {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		} */

		w.WriteHeader(http.StatusCreated)

	}
}
