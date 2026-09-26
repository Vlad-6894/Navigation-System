package log_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *LogService) CheckUnique(
	ctx context.Context,
	uuid uuid.UUID,
) (bool, error) {
	ok, err := s.cache.CheckUnique(ctx, uuid)
	if err != nil {
		return false, fmt.Errorf("fail check unique from repository: %w", err)
	}

	return ok, nil
}
