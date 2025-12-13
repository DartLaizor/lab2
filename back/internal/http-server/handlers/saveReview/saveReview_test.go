package savereview_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	savereview "server/internal/http-server/handlers/saveReview"
	"server/internal/http-server/handlers/saveReview/mocks"
	"server/internal/storage"
)

func TestSaveRewiev(t *testing.T) {
	cases := []struct {
		name          string
		inputReview   storage.Review
		mockError     error
		wantRespError string
		wantCode      int
	}{
		{
			name: "Success",
			inputReview: storage.Review{
				Name:         "Анна Иванова",
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "anna@example.com",
				Technologies: []string{"Go", "PostgreSQL"},
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			wantCode: http.StatusCreated,
		},
		{
			name: "Missing name",
			inputReview: storage.Review{
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "anna@example.com",
				Technologies: []string{"Go"},
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			wantRespError: "Invalid JSON",
			wantCode:      http.StatusBadRequest,
		},
		{
			name: "Invalid email",
			inputReview: storage.Review{
				Name:         "Анна Иванова",
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "invalid-email",
				Technologies: []string{"Go"},
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			wantRespError: "Invalid JSON",
			wantCode:      http.StatusBadRequest,
		},
		{
			name: "Empty technologies",
			inputReview: storage.Review{
				Name:         "Анна Иванова",
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "anna@example.com",
				Technologies: []string{}, // нарушает validate:"min=1"
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			wantRespError: "Invalid JSON",
			wantCode:      http.StatusBadRequest,
		},
		{
			name: "Wrong HTTP method",
			inputReview: storage.Review{
				Name:         "Анна Иванова",
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "anna@example.com",
				Technologies: []string{"Go"},
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			// Но запрос будет GET → проверим отдельно
			wantRespError: "Use method POST",
			wantCode:      http.StatusMethodNotAllowed,
		},
		{
			name: "DB Save Error",
			inputReview: storage.Review{
				Name:         "Анна Иванова",
				Date:         "2025-12-08",
				Phone:        "+79001234567",
				Email:        "anna@example.com",
				Technologies: []string{"Go"},
				Rating:       4,
				Comment:      "Хорошо работает!",
			},
			mockError:     errors.New("DB connection failed"),
			wantRespError: "couldn't save to the DataBase",
			wantCode:      http.StatusInternalServerError,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockDB := mocks.NewReviewsSaver(t)

			logger := slog.Default()

			handler := savereview.SaveReviewsHadnler(logger, mockDB)

			var req *http.Request

			if tt.name == "Wrong HTTP method" {
				req = httptest.NewRequest(http.MethodGet, "/review", nil)
			} else {
				// Все остальные — POST с телом
				body, _ := json.Marshal(tt.inputReview)
				req = httptest.NewRequest(http.MethodPost, "/review", bytes.NewReader(body))
				req.Header.Set("Content-Type", "application/json")

				// Настраиваем мок, если ожидается вызов SaveReview
				if tt.wantCode != http.StatusBadRequest && tt.wantCode != http.StatusMethodNotAllowed {
					mockDB.On("SaveReview", mock.MatchedBy(func(r storage.Review) bool {
						return r.Name == tt.inputReview.Name &&
							r.Email == tt.inputReview.Email
					})).Return(tt.mockError).Once()
				}
			}

			rr := httptest.NewRecorder()
			handler(rr, req)

			require.Equal(t, tt.wantCode, rr.Code)

			// Проверяем тело ответа только если ожидается ошибка в теле
			if tt.wantRespError != "" && tt.wantCode != http.StatusMethodNotAllowed {
				// В текущей реализации хендлер не возвращает JSON — только текст ошибки
				require.Contains(t, rr.Body.String(), tt.wantRespError)
			} else if tt.wantRespError != "" && tt.wantCode == http.StatusMethodNotAllowed {
				require.Contains(t, rr.Body.String(), tt.wantRespError)
			}
			// При успехе тело пустое — только статус 201

			mockDB.AssertExpectations(t)
		})

	}

}
