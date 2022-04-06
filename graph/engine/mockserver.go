package engine

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/speedoops/go-gqlrest-demo/config"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/tal-tech/go-zero/core/logx"
)

var _server http.Handler

func GetMockServer(resolver generated.ResolverRoot) http.Handler {
	if _server != nil {
		return _server
	}

	return NewMockServer(resolver)
}

func NewMockServer(resolver generated.ResolverRoot) http.Handler {
	// 1. 初始化服务端配置
	var c config.Config
	//conf.MustLoad(FindConfigFile("config.yaml"), &c)
	config.GraphQL = c.GraphQL
	c.Log.Mode = "console"
	logx.MustSetup(c.Log)

	// 2. 运行 GraphQL Server
	srv := NewServer(resolver)

	mux := chi.NewRouter()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	generated.RegisterHandlers(mux, srv, "")

	return mux
}

func FindConfigFileByName(fileName string) string {
	_, progName, _, _ := runtime.Caller(0)
	lastDir := path.Dir(progName)
	for {
		tryPath := filepath.Join(lastDir, fileName)
		if fi, err := os.Stat(tryPath); err == nil {
			if mode := fi.Mode(); mode.IsRegular() {
				return tryPath
			}
		}

		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			panic(fmt.Sprintf("config file '%s' not found", fileName))
		}
		lastDir = newDir
	}
}
