package secondary

import (
	"context"
)

type SecondaryAdapter struct{}

func NewSecondaryAdapter() (*SecondaryAdapter, error) {
	return &SecondaryAdapter{}, nil
}

func (a *SecondaryAdapter) Run(ctx context.Context) error {
	return nil
}
