package domain

import (
	"context"

	"github.com/nxdir-s/go-hexarch/internal/ports"
)

type ErrRunDomain struct {
	err error
}

func (e *ErrRunDomain) Error() string {
	return "failed to run orchestrator: " + e.err.Error()
}

type Orchestrator struct {
	service ports.Service
}

func NewOrchestrator(service ports.Service) *Orchestrator {
	return &Orchestrator{
		service: service,
	}
}

func (d *Orchestrator) Run(ctx context.Context) error {
	if err := d.service.Update(ctx); err != nil {
		return &ErrRunDomain{err}
	}

	return nil
}
