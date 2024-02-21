package main

import (
	"hexagonal-architecture-go/internal/adapters/app/api"
	"hexagonal-architecture-go/internal/adapters/core/arithmetic"
	"hexagonal-architecture-go/internal/adapters/framework/right_driven/db"
	"hexagonal-architecture-go/internal/ports"
	"log"
	"os"

	// "hexagonal-architecture-go/internal/adapters/app/api"
	// "hexagonal-architecture-go/internal/adapters/core/arithmetic"
	// "hexagonal-architecture-go/internal/adapters/framework/right_driven/db"
	gRPC "hexagonal-architecture-go/internal/adapters/framework/left_driver/grpc"
	// "hexagonal-architecture-go/internal/ports"
)

func main() {

	var err error

	// Ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection %v", err)
	}

	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}
