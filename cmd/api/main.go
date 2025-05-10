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
	var database ports.Database

	// services
	var coreService ports.Service

	// domain orchestrators
	var orchestrator ports.Orchestrator

	// primary adapter
	var adapter ports.API

	database = secondary.NewDatabaseAdapter()

	coreService = service.NewService(database)

	orchestrator = domain.NewOrchestrator(coreService)

	adapter = primary.NewAPIAdapter(orchestrator)

	if err := adapter.Run(ctx); err != nil {
		fmt.Fprintf(os.Stdout, "error running adapter: %s\n", err.Error())
		os.Exit(1)
	}
}
