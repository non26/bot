package service

import (
	"context"
)

func (s *botTemplateService) Delete(ctx context.Context, id string) error {
	return s.botTemplateRepository.Delete(ctx, id)
}
