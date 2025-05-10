package primary

import (
	"context"

	"github.com/nxdir-s/go-hexarch/internal/ports"
)

type APIAdapter struct {
	domain ports.Orchestrator
}

func NewAPIAdapter(domain ports.Orchestrator) *APIAdapter {
	return &APIAdapter{
		domain: domain,
	}
}

func (a *APIAdapter) Run(ctx context.Context) error {
	return a.domain.Run(ctx)
}
