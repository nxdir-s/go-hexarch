package service

import (
	"context"

	"github.com/nxdir-s/go-hexarch/internal/ports"
)

type Service struct {
	adapter ports.Database
}

func NewService(adapter ports.Database) *Service {
	return &Service{
		adapter: adapter,
	}
}

func (s *Service) Update(ctx context.Context) error {
	return s.adapter.Update(ctx)
}
