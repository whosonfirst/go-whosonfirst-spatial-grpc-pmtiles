package server

import (
	"flag"
	"fmt"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

func DefaultFlagSet() (*flag.FlagSet, error) {

	fs, err := spatial_flags.CommonFlags()

	if err != nil {
		return nil, fmt.Errorf("Failed to create common flags, %v", err)
	}

	err = spatial_flags.AppendIndexingFlags(fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to append indexing flags, %v", err)
	}

	err = grpc_flags.AppendGRPCServerFlags(fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to append server flags, %v", err)
	}

	return fs, nil
}
