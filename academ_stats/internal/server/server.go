package server

import (
	"academ_stats/internal/api"
	"academ_stats/internal/config"
	"academ_stats/internal/repository/pb/excel_table"
	"academ_stats/internal/repository/pb/session_manager"
	"academ_stats/internal/repository/postgres"
	"academ_stats/internal/service"
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

type Server interface {
	Run(ctx context.Context)
	Stop(ctx context.Context)
}

type server struct {
	router         *echo.Echo       //
	sessionManager *grpc.ClientConn // send request to service "session_manager"
	excelTable     *grpc.ClientConn // send request to service "excel_table"
	cfg            *config.Config   //
}

func NewServer(env *Env, cfg *config.Config) Server {
	s := server{
		router:         echo.New(),
		sessionManager: newGrpcCli(cfg.SessionAddr),
		excelTable:     newGrpcCli(cfg.ExcelAddr),
		cfg:            cfg,
	}

	// storage
	storage := postgres.NewStorage(env.pool)

	// service
	grpcCli := service.GrpcCli{
		SessionManager: session_manager.NewSessionManagerClient(s.sessionManager),
		ExcelTable:     excel_table.NewExcelTableClient(s.excelTable),
	}
	svc := service.New(storage, &grpcCli, env.zoApi, cfg.Debug)

	// handlers
	hndl := api.NewHandlers(s.router.Logger, svc, cfg.Debug)

	// set middlewares
	s.router.Use(middleware.Logger(), middleware.Recover())

	// register handlers
	g := s.router.Group("/api/academ-stats")

	g.GET("/module-list", hndl.ModuleList)
	// g.GET("/cadet-time", hndl.GetComputers)

	// journey
	gj := g.Group("/journey")
	gj.GET("/top-cadets", hndl.TopCadets)          // /top-cadets?id=<moduleID>
	gj.GET("/top-cadets-file", hndl.TopCadetsFile) // /top-cadets?id=<moduleID>

	// // audits
	// ga := g.Group("/audits")
	// ga.GET("/top-cadets", hndl.TopCadets)          // /top-cadets?id=<moduleID>
	// ga.GET("/top-cadets-file", hndl.TopCadetsFile) // /top-cadets?id=<moduleID>

	// other
	gc := g.Group("/check")
	gc.GET("/token-expire-date", hndl.GetTokenExpDate)

	return &s
}

func (s *server) Run(ctx context.Context) {
	ctxSignal, cancelSignal := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// start server
	go func() {
		defer cancelSignal()

		if err := s.router.Start(s.cfg.APIPort); err != nil && err != http.ErrServerClosed {
			log.Printf("server start error: %s\n", err.Error())
		}
	}()

	// wait system notifiers or cancel func
	<-ctxSignal.Done()
}

func (s *server) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.sessionManager.Close(); err != nil {
		log.Printf("grpc \"session_manager\" cli server stop error: %s\n", err.Error())
		return
	}

	if err := s.excelTable.Close(); err != nil {
		log.Printf("grpc \"excel_table\" cli server stop error: %s\n", err.Error())
		return
	}

	if err := s.router.Shutdown(ctx); err != nil {
		log.Printf("server stop error: %s\n", err.Error())
		return
	}
	log.Println("server stopped successfully")
}
