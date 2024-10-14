package main

import (
	"academ_stats/internal/config"
	"academ_stats/internal/server"
	"context"
)

func main() {
	ctx := context.Background()

	// read conf
	cfg := config.Configure()
	cfg.Print()

	// envorinments [db and etc...]
	env := server.NewEnv(ctx, cfg)
	defer env.Stop(ctx)

	// server
	srv := server.NewServer(env, cfg)
	srv.Run(ctx)
	defer srv.Stop(ctx)
}
