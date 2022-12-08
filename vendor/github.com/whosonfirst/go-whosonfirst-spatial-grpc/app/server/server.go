package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/lookup"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	grpc_server "github.com/whosonfirst/go-whosonfirst-spatial-grpc/server"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	spatial_app "github.com/whosonfirst/go-whosonfirst-spatial/app"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run(ctx context.Context, logger *log.Logger) error {

	fs, err := DefaultFlagSet()

	if err != nil {
		return fmt.Errorf("Failed to derive default flag set, %w", err)
	}

	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	err := spatial_flags.ValidateCommonFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate common flags, %v", err)
	}

	err = spatial_flags.ValidateIndexingFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate indexing flags, %v", err)
	}

	err = grpc_flags.ValidateGRPCServerFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate server flags, %v", err)
	}

	host, _ := lookup.StringVar(fs, grpc_flags.HOST)
	port, _ := lookup.IntVar(fs, grpc_flags.PORT)

	sp_app, err := spatial_app.NewSpatialApplicationWithFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to create new spatial application, %v", err)
	}

	uris := fs.Args()

	err = sp_app.IndexPaths(ctx, uris...)

	if err != nil {
		return fmt.Errorf("Failed to index paths, %v", err)
	}

	spatial_server, err := grpc_server.NewSpatialServer(sp_app)

	if err != nil {
		return fmt.Errorf("Failed to create spatial server, %v", err)
	}

	grpc_server := grpc.NewServer()

	spatial.RegisterSpatialServer(grpc_server, spatial_server)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening on %s\n", addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpc_server.Serve(lis)
	return nil
}
