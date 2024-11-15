package db

import (
	"go.uber.org/zap"
)

type options struct {
	cache  bool
	logger *zap.Logger
}

type Option interface {
	apply(*options)
}

func Open(
	addr string,
	opts ...Option,
) {

}
