package ports

import "context"

type Database interface {
	Update(ctx context.Context) error
}
