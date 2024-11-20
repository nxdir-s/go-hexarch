package ports

import "context"

type PrimaryPort interface {
	Run(ctx context.Context) error
}
