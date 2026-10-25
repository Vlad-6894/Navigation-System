package pkg_grpc_server

import (
	"context"
	"fmt"
	pkg_logger "navigation-pkg/logger"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	cfg    GRPCServerConfig
	log    *pkg_logger.Logger
	Server *grpc.Server
}

func NewGRPCServer(
	cfg GRPCServerConfig,
	log *pkg_logger.Logger,
) *GRPCServer {
	return &GRPCServer{
		cfg:    cfg,
		log:    log,
		Server: grpc.NewServer(),
	}
}

func (s *GRPCServer) Run(ctx context.Context) error {
	errorChannel := make(chan error, 1)

	listen, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		return fmt.Errorf("fail to get listener: %w", err)
	}

	go func() {
		defer close(errorChannel)

		if err := s.Server.Serve(listen); err != nil {
			errorChannel <- fmt.Errorf("fail to start grpc server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.Server.GracefulStop()

	case err := <-errorChannel:
		if err != nil {
			return fmt.Errorf("grpc server error: %w", err)
		}
	}

	return nil
}
