package main

import (
	"context"
	"excel_table/config"
	"excel_table/internal/server"
)

func main() {
	ctx := context.Background()

	// read conf
	cfg := config.Configure()

	// envorinments [db and etc...]
	env := server.NewEnv(ctx, cfg)
	defer env.Stop(ctx)

	// server
	srv := server.NewServer(env, cfg)
	srv.Run(ctx)
	defer srv.Stop(ctx)
}
