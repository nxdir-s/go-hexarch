package ports

import "context"

type Orchestrator interface {
	Run(ctx context.Context) error
}

type Service interface {
	Run(ctx context.Context) error
}
