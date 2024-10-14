package server

import (
	"net"

	"context"
	"excel_table/config"
	"excel_table/internal/repository/pb/excel_table"
	"excel_table/internal/repository/postgres"
	"excel_table/internal/service"
	"log"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type Server interface {
	Run(ctx context.Context)
	Stop(ctx context.Context)
}

type server struct {
	grpcCli *grpc.Server   //
	cfg     *config.Config //
}

func NewServer(env *Env, cfg *config.Config) Server {
	s := server{
		grpcCli: grpc.NewServer(),
		cfg:     cfg,
	}
	// storage
	storage := postgres.NewStorage(env.pool)

	// service for grpc
	svcGrpc := service.NewServiceGrpc(storage)
	excel_table.RegisterExcelTableServer(s.grpcCli, svcGrpc)

	return &s
}

func (s *server) Run(ctx context.Context) {
	ctxSignal, cancelSignal := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// start grpc server
	go func() {
		defer cancelSignal()

		log.Printf("grpc starts on port: %s\n", s.cfg.GRPCPort)

		lis, err := net.Listen("tcp", s.cfg.GRPCPort)
		if err != nil {
			log.Printf("net listen error: %s\n", err.Error())
			return
		}

		if err := s.grpcCli.Serve(lis); err != nil {
			log.Printf("server start error: %s\n", err.Error())
		}
	}()

	// wait system notifiers or cancel func
	<-ctxSignal.Done()
}

func (s *server) Stop(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	s.grpcCli.GracefulStop()

	log.Println("server stop done")
}
