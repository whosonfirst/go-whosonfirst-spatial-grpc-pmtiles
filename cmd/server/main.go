package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-spatial-pmtiles"
)

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/app/server"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run client, %v", err)
	}
}
