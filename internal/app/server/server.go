package server

import (
	"GraphQL-project/internal/app/server/graphqlServer"
	"GraphQL-project/internal/app/server/grpcServer"
	"GraphQL-project/internal/config"
	"GraphQL-project/internal/handlers/grpcHandlers"
	"GraphQL-project/internal/scenarios"
	"GraphQL-project/internal/storage"
	"GraphQL-project/internal/storage/mongoDb"
	"context"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
	"log"
)

type server struct {
	config    *config.Config
	scenarios scenarios.User
	storage   storage.Storage
	ctx       context.Context
}

func New() *server {
	ctx := context.Background()

	storage, err := mongoDb.New(ctx)
	if err != nil {
		log.Fatalln("error in connection mongo : %w", err)
	}
	return &server{
		config:    config.New(),
		scenarios: scenarios.New(storage),
		storage:   storage,
		ctx:       ctx,
	}
}

func (s *server) Start() error {
	gr, ctx := errgroup.WithContext(s.ctx)

	gr.Go(func() error {
		return grpcServer.New(grpcHandlers.New(s.scenarios), s.config).Start(ctx)
	})
	gr.Go(func() error {
		return graphqlServer.New(s.scenarios, s.config).Start(ctx)
	})
	if err := gr.Wait(); err != nil {
		slog.Error(err.Error())
	}
	return nil
}
