package main

import (
	// application
	"github.com/sunil206b/hex-arc-app/internal/application/api"
	"github.com/sunil206b/hex-arc-app/internal/application/core/arithmetic"

	// Adapters
	gRPC "github.com/sunil206b/hex-arc-app/internal/adapters/framework/left/grpc"
	"github.com/sunil206b/hex-arc-app/internal/adapters/framework/right/db"

	"os"
)

func main() {

	dbDriver := os.Getenv("DB_DRIVER")
	dsName := os.Getenv("DB_NAME")
	dbAdapter := db.NewAdapter(dbDriver, dsName)
	defer dbAdapter.CloseDbConnection()

	//Core
	core := arithmetic.New()
	applicationAPI := api.NewApplication(dbAdapter, core)

	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()
}
