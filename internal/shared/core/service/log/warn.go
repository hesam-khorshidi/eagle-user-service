package log

import "context"

func (s *Service) Warn(ctx context.Context, msg string, args ...any) {
	s.logger.WarnContext(ctx, msg, args...)
}
