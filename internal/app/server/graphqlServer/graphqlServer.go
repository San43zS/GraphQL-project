package graphqlServer

import (
	"GraphQL-project/graph"
	"GraphQL-project/internal/config"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"golang.org/x/exp/slog"
	"log"
	"net/http"

	"context"
)

type Server struct {
	config    *config.Config
	scenarios graph.UserList
}

func New(scenarios graph.UserList, config *config.Config) *Server {
	return &Server{
		config:    config,
		scenarios: scenarios,
	}
}

func (s *Server) Start(ctx context.Context) (err error) {
	resolver := graph.Resolver{
		UserList: s.scenarios,
	}
	graphqlSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	http.Handle("/graph", playground.Handler("Graphql playground", "/list"))
	http.Handle("/list", graphqlSrv)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.GraphqlPort),
		Handler: nil,
	}
	log.Printf("GraphQL server is runnning on http://localhost:%d/graph \n", s.config.GraphqlPort)

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- srv.ListenAndServe()
	}()
	select {
	case err = <-errCh:
	case <-ctx.Done():
		err = srv.Shutdown(ctx)
		slog.Error(err.Error())
	}
	return err
}
