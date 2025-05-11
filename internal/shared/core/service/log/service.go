package log

import (
	"log/slog"
	"os"

	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
)

var _ inbound.LogService = (*Service)(nil)

type Service struct {
	logger *slog.Logger
}

func New(level slog.Level) *Service {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(handler)
	return &Service{logger: logger}
}
