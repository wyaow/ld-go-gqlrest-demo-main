package engine

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/speedoops/go-gqlrest-demo/graph/errorsx"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/speedoops/go-gqlrest-demo/graph/model"
	"github.com/speedoops/go-gqlrest/handlerx"
	"github.com/tal-tech/go-zero/core/logx"
)

func NewServer(resolver generated.ResolverRoot) *handler.Server {
	c := generated.Config{Resolvers: resolver}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		if !getCurrentUser(ctx).HasRole(role) {
			// block calling the next resolver
			return nil, errorsx.NewNotAllowedError(errors.New("access denied"))
		}
		log.Println("hasRole")

		// or let it pass through
		return next(ctx)
	}

	c.Directives.Hide = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		forArg []string) (res interface{}, err error) {
		return next(ctx)
	}
	c.Directives.Http = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		url string, method *string) (res interface{}, err error) {
		return next(ctx)
	}
	c.Directives.Preview = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		toggledBy string) (res interface{}, err error) {
		return next(ctx)
	}

	srv := handlerx.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.SetErrorPresenter(errorsx.AppErrorPresenter)
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logx.ErrorStack("internal server error")
		return errors.New("internal server error")
	})

	return srv
}

func getCurrentUser(ctx context.Context) model.User {
	return model.User{}
}
