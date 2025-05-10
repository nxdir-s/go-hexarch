package ports

import "context"

type API interface {
	Run(ctx context.Context) error
}
