package main

import (
	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg"
)

type receiver struct {
	logger *zerolog.Logger
}

func (r *receiver) BindLogger(logger *zerolog.Logger) {
	r.logger = logger
}

func (r *receiver) Logger() *zerolog.Logger {
	return r.logger
}

func (r *receiver) Init(rt pkg.IsRuntime) {
	rt.Events().On("test", func(args ...any) {
		r.logger.Info().Msgf("event 'test' recieved, args are: %+v", args)
	})
}

func (r *receiver) Name() string {
	return "receiver"
}
