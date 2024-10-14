package service

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/request"
	"academ_stats/internal/domain/response"
	"academ_stats/internal/helper/parser"
	"academ_stats/internal/repository/pb/excel_table"
	"academ_stats/internal/repository/pb/session_manager"
	"academ_stats/internal/repository/postgres"
	"academ_stats/internal/repository/zero_one_api"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const AcademieStorageApiUrl = "https://zero.academie.one/api/storage"

type Service interface {
	TopCadets(ctx context.Context, pathID int) (*response.TopCadetsResponse, error)
	TopCadetsFile(ctx context.Context, pathID int) ([]byte, error)
	GetTokenExpDate(ctx context.Context) (*response.TokenExpireDate, error)
	ModuleList(ctx context.Context) ([]domain.Module, error)
}

func New(storage postgres.Storage, grpcCli *GrpcCli, zoApi zero_one_api.ZeroOneApi, debug bool) Service {
	return &service{
		storage: storage,
		grpcCli: grpcCli,
		zoApi:   zoApi,
		debug:   debug,
	}
}

type service struct {
	storage postgres.Storage
	zoApi   zero_one_api.ZeroOneApi
	grpcCli *GrpcCli
	debug   bool
}

type GrpcCli struct {
	SessionManager session_manager.SessionManagerClient
	ExcelTable     excel_table.ExcelTableClient
}

func (s *service) TopCadets(ctx context.Context, pathID int) (*response.TopCadetsResponse, error) {
	pr := domain.Period{
		FromDate: time.Now().AddDate(0, -6, 0), // minus six month
		ToDate:   time.Now(),
	}

	// parse & sort data
	tc, err := s.topCadetsGeneral(ctx, pathID, &pr)
	if err != nil {
		return nil, fmt.Errorf("general: %s", err)
	}

	// no data
	if len(tc.Cadets) == 0 {
		return &response.TopCadetsResponse{
			Current:  []byte("[]"),
			Last:     []byte("[]"),
			LastDate: time.Unix(0, 0),
		}, response.ErrNotFound
	}

	// copy data to response struct &
	// prepare current response for db
	curr, err := json.Marshal(parser.TopCadetsRest(tc))
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	// write to db
	return s.storage.TopCadets(ctx, curr)
}

func (s *service) TopCadetsFile(ctx context.Context, pathID int) ([]byte, error) {
	pr := domain.Period{
		FromDate: time.Now().AddDate(0, -6, 0), // minus six month
		ToDate:   time.Now(),
	}

	// parse & sort data
	tc, err := s.topCadetsGeneral(ctx, pathID, &pr)
	if err != nil {
		return nil, fmt.Errorf("general: %s", err)
	}

	// if no data
	if len(tc.Cadets) == 0 {
		return nil, nil
	}

	// get excel table
	res, err := s.grpcCli.ExcelTable.GetTopCadets(ctx, parser.TopCadetsGrpc(tc))
	if err != nil {
		return nil, fmt.Errorf("grpc: excel table: top cadets: %s", err)
	}

	return res.File, nil
}

func (s *service) GetTokenExpDate(ctx context.Context) (*response.TokenExpireDate, error) {
	expDate, _, err := s.zoApi.Token()
	if err != nil {
		return nil, fmt.Errorf("token: %w", err)
	}

	return &response.TokenExpireDate{
		Expire:    time.Unix(expDate, 0),
		PreNotify: time.Unix(expDate-int64(zero_one_api.RefreshTokenPreNotify.Seconds()), 0),
	}, nil
}

func (s *service) topCadetsGeneral(ctx context.Context, pathID int, pr *domain.Period) (*domain.TopCadets, error) {
	wg := sync.WaitGroup{}
	chErr := make(chan error, 1)

	var err error
	var src *request.TopCadets
	var hours map[int]domain.HoursDTO
	var adm map[int]domain.Piscine

	// get cadets journey {module, checkpoint, piscine}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var e error
		src, e = s.topCadetsGraphql(ctx, pathID)
		if e != nil {
			chErr <- e
		}
	}()

	// get time management
	wg.Add(1)
	go func() {
		defer wg.Done()
		hours = s.hoursGrpc(ctx, pathID, pr)
	}()

	// get admissions
	wg.Add(1)
	go func() {
		defer wg.Done()
		adm = s.admissionDB(ctx)
	}()

	// waiting for all the goroutines to be finalized
	wg.Wait()
	close(chErr)

	// check error
	if err = <-chErr; err != nil {
		return nil, fmt.Errorf("graphql: %s", err)
	}

	if s.debug {
		log.Printf("all services finished: is nil: {src: %t, hours: %t, adm: %t}\n",
			src == nil, hours == nil, adm == nil)

		log.Printf("src length is %d\n", len(src.Cadets))
	}

	// if no data
	if len(src.Cadets) == 0 {
		return nil, nil
	}

	// parse & sort data
	return parser.TopCadetsGraphql(src, adm, hours, pr), nil
}

func (s *service) topCadetsGraphql(ctx context.Context, pathID int) (*request.TopCadets, error) {
	if s.debug {
		log.Println("top cadets graphql service started")
		defer log.Println("top cadets graphql service finished")
	}

	// send graphql request
	resp, err := s.zoApi.TopCadets(pathID)
	if err != nil {
		return nil, fmt.Errorf("zero one: top cadets: %w", err)
	}

	// parse data
	tcReq := request.TopCadets{}

	err = json.Unmarshal(resp, &tcReq)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return &tcReq, nil
}

func (s *service) hoursGrpc(ctx context.Context, pathID int, pr *domain.Period) map[int]domain.HoursDTO {
	if s.debug {
		log.Println("hours grpc service started")
		defer log.Println("hours grpc service finished")
	}

	totalTimes, err := s.grpcCli.SessionManager.GetCadetsTimeByModuleID(ctx, &session_manager.CadetsTimeRequest{
		ModuleId: int32(pathID),
		FromDate: timestamppb.New(pr.FromDate),
		ToDate:   timestamppb.New(pr.ToDate),
	})
	if err != nil {
		log.Printf("grpc: cadets hours: %s", err) // not critical
		return nil
	}

	return parser.HoursGrpc(totalTimes)
}

func (s *service) admissionDB(ctx context.Context) map[int]domain.Piscine {
	if s.debug {
		log.Println("admission db service started")
		defer log.Println("admission db service finished")
	}

	adm, err := s.storage.PiscineList(ctx)
	if err != nil {
		log.Printf("db: did not find any admissions: %s", err) // not critical
		return nil
	}

	return adm
}

func (s *service) ModuleList(ctx context.Context) ([]domain.Module, error) {
	return s.storage.ModuleList(ctx)
}
