package internal

// это возможно obsolete

import (
	"log/slog"
	"os"
)

func SetupVerboseLogging() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
	slog.Info("Verbose logging enabled")
}

func SetupBasicLogging() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))
}
