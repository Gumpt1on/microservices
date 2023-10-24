package main

import (
	"github.com/Gumpt1on/microservices/order/config"
	"github.com/Gumpt1on/microservices/order/internal/adapters/db"
	"github.com/Gumpt1on/microservices/order/internal/adapters/grpc"
	"github.com/Gumpt1on/microservices/order/internal/application/core/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}