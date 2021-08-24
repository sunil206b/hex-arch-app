package main

import (
	"github.com/sunil206b/hex-arc-app/internal/adapters/app/api"
	"github.com/sunil206b/hex-arc-app/internal/adapters/core/arithmetic"
	gRPC "github.com/sunil206b/hex-arc-app/internal/adapters/framework/left/grpc"
	"github.com/sunil206b/hex-arc-app/internal/adapters/framework/right/db"
	"github.com/sunil206b/hex-arc-app/internal/ports"
	"os"
)

func main() {
	//Ports
	var core ports.ArithmeticPort
	var dbAdapter ports.DBPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dsName := os.Getenv("DB_NAME")
	dbAdapter = db.NewAdapter(dbDriver, dsName)
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
