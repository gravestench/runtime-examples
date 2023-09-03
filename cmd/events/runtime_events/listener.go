package main

import (
	"github.com/gravestench/runtime"
	"github.com/rs/zerolog"
)

type listensForNewServices struct {
	logger *zerolog.Logger
}

func (s *listensForNewServices) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *listensForNewServices) Logger() *zerolog.Logger {
	return s.logger
}

func (s *listensForNewServices) Init(rt runtime.Runtime) {
	// noop
}

func (s *listensForNewServices) Name() string {
	return "listener"
}

// there are a bunch of runtime events to bind to via
// implementing an interface like this one

func (s *listensForNewServices) OnServiceAdded(args ...interface{}) {
	if len(args) < 1 {
		return
	}

	service, ok := args[0].(runtime.Service)
	if !ok {
		return
	}

	if service == s {
		return
	}

	s.logger.Info().Msgf("found another service %q", service.Name())
}
