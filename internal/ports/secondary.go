package ports

import "context"

type SecondaryPort interface {
	Run(ctx context.Context) error
}
