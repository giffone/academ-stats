STATS_DATABASE_URL ?= postgres://user:password@host:port/db-name?sslmode=disable
EXCEL_DATABASE_URL ?= postgres://user:password@host:port/db-name?sslmode=disable

.PHONY: migrate build run run_local down gen_proto

gen_proto:
	protoc -I=./academ_stats/pkg/proto --go_out=./academ_stats --go-grpc_out=./academ_stats session_manager.proto excel_table.proto
	protoc -I=./excel_table/pkg/proto --go_out=./excel_table --go-grpc_out=./excel_table excel_table.proto

build:
	docker compose build

run: build
	docker compose up -d

migrate: run
	docker compose exec academ_stats migrate -path db/migrations -database "$(STATS_DATABASE_URL)" up
	docker compose exec excel_table migrate -path db/migrations -database "$(EXCEL_DATABASE_URL)" up

down:
	docker compose down

run_local:
	@echo "Define local run configuration if needed"
