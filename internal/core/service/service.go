package service

import (
	"context"

	"github.com/nxdir-s/go-hexarch/internal/ports"
)

type Service struct {
	adapter ports.SecondaryPort
}

func NewService(adapter ports.SecondaryPort) (*Service, error) {
	return &Service{
		adapter: adapter,
	}, nil
}

func (s *Service) Run(ctx context.Context) error {
	return s.adapter.Run(ctx)
}
