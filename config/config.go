package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	GraphQL GraphQLConf
}

type GraphQLConf struct {
	Debug struct {
		EnableVerbose bool `json:",optional"` //nolint:revive,staticcheck // FIXME: go-zero
	}
}

// TODO: 破坏了代码分层，后续优化
var GraphQL GraphQLConf
