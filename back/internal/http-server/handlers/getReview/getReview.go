package savereview

import (
	"encoding/json"
	"net/http"

	"server/internal/storage"
)

type ReviewsSaver interface {
	SaveRewiev(r storage.Review) error
}

func SaveReviewsHadnler(db ReviewsSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Use method POST ", http.StatusMethodNotAllowed)
			return
		}

		var review storage.Review

		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if review.Name == "" || review.Email!="" || review.Phone != "" || len(review.Technologies)!=0 || review.Date == "" || review.Rating < 0 || review.Rating > 9 {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
