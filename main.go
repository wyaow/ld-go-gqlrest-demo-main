package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/google/gops/agent"
	"github.com/speedoops/go-gqlrest-demo/graph"
	"github.com/speedoops/go-gqlrest-demo/graph/engine"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/speedoops/go-gqlrest/handlerx"
)

const defaultPort = "8080"

func main() {
	if err := agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// var c config.Config
	// //conf.MustLoad(FindConfigFile("config.yaml"), &c)
	// config.GraphQL = c.GraphQL
	// c.Log.Mode = "console"
	// logx.MustSetup(c.Log)

	srv := engine.NewServer(&graph.Resolver{})

	mux := chi.NewRouter()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	mux.Handle("/graphql", srv)
	generated.RegisterHandlers(mux, srv, "")
	handlerx.RegisterPrinter(&LogPrinter{})

	_ = mux.ServeHTTP // 调试入口1：HTTP Server Entry
	_ = srv.ServeHTTP // 调试入口2：GraphQL Handler Entry
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

type LogPrinter struct{}

func (l *LogPrinter) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *LogPrinter) Println(v ...interface{}) {
	log.Println(v...)
}
