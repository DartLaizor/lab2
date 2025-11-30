package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"server/internal/config"
	savereview "server/internal/http-server/handlers/saveReview"
	sendreview "server/internal/http-server/handlers/sendReview"
	mwLogger "server/internal/http-server/middleware/logger"
	"server/internal/logger"
	"server/internal/logger/sl"
	"server/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.NewLogger(cfg.Env)
	_ = log

	db, err := postgres.NewStorage(cfg.StoragePath)

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()
	//*middlewares
	router.Use(middleware.Recoverer)

	router.Use(mwLogger.New(log))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//*handlers

	router.Get("/send", sendreview.SendReviewsHandler(log,db))
	router.Post("/save", savereview.SaveReviewsHadnler(log,db))
	//TODO: run server
	log.Info("starting server", slog.String("address", cfg.Address))

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start server ")
		}
	}()

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	db.CloseStorage()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))

		return
	}

	// TODO: close storage

	log.Info("server stopped")
}
