package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// NewLogger возвращает настроенный экземпляр slog.Logger в зависимости от среды.
func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "local":
		// Для локальной разработки используем tint для цветного и удобочитаемого вывода
		logger = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				Level:      slog.LevelDebug, // Уровень логирования (DEBUG и выше)
				TimeFormat: time.Kitchen,    // Формат времени, например "3:04PM"
				// NoColor: false, // Цвета включены по умолчанию
			}),
		)
	case "prod":
		// Для продакшена используем JSON формат для лучшей обработки логов
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo, // Уровень логирования (INFO и выше)
			}),
		)
	default:
		// По умолчанию (например, для "dev", "test") можно использовать JSON или Text
		// Здесь для примера используем JSON
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}

	return logger
}
