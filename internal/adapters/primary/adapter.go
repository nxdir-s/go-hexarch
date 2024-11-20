package primary

import (
	"context"

	"github.com/nxdir-s/go-hexarch/internal/ports"
)

type PrimaryAdapter struct {
	domain ports.Orchestrator
}

func NewPrimaryAdapter(domain ports.Orchestrator) (*PrimaryAdapter, error) {
	return &PrimaryAdapter{
		domain: domain,
	}, nil
}

func (a *PrimaryAdapter) Run(ctx context.Context) error {
	return a.domain.Run(ctx)
}
