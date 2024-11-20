package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/nxdir-s/go-hexarch/internal/adapters/primary"
	"github.com/nxdir-s/go-hexarch/internal/adapters/secondary"
	"github.com/nxdir-s/go-hexarch/internal/core/domain"
	"github.com/nxdir-s/go-hexarch/internal/core/service"
	"github.com/nxdir-s/go-hexarch/internal/ports"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// secondary adapters
	var secondaryAdapter ports.SecondaryPort

	// services
	var coreService ports.Service

	// domain orchestrators
	var orchestrator ports.Orchestrator

	// primary adapter
	var adapter ports.PrimaryPort

	secondaryAdapter, err := secondary.NewSecondaryAdapter()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating secondary adapter: %s\n", err.Error())
		os.Exit(1)
	}

	coreService, err = service.NewService(secondaryAdapter)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating service: %s\n", err.Error())
		os.Exit(1)
	}

	orchestrator, err = domain.NewOrchestrator(coreService)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating domain orchestrator: %s\n", err.Error())
		os.Exit(1)
	}

	adapter, err = primary.NewPrimaryAdapter(orchestrator)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error creating primary adapter: %s\n", err.Error())
		os.Exit(1)
	}

	if err := adapter.Run(ctx); err != nil {
		fmt.Fprintf(os.Stdout, "error running adapter: %s\n", err.Error())
		os.Exit(1)
	}
}
