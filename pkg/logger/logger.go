package logger

import (
	"log/slog"
	"os"
)

func InitLogger() {
	// SLogの設定 JSON形式で出力
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
