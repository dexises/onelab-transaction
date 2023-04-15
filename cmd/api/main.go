package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"onelab/internal/config"
	"onelab/internal/jsonlog"
	"onelab/internal/repository"
	"onelab/internal/service"
	"onelab/internal/transport/http"
	"onelab/internal/transport/http/handler"

)

func main() {
	log.Fatalf(fmt.Sprintf("Service shut down %v", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	repo, err := repository.NewRepository(ctx, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := service.NewService(cfg, repo)

	h := handler.NewManager(cfg, srv)

	server := http.NewServer(cfg, h)

	return server.StartServer(ctx, logger)
}
