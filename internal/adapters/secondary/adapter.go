package secondary

import (
	"context"
)

type DatabaseAdapter struct{}

func NewDatabaseAdapter() *DatabaseAdapter {
	return &DatabaseAdapter{}
}

func (a *DatabaseAdapter) Update(ctx context.Context) error {
	return nil
}
