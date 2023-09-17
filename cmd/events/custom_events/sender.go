package main

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/runtime/pkg"
)

type sender struct {
	logger *zerolog.Logger
}

func (s *sender) BindLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func (s *sender) Logger() *zerolog.Logger {
	return s.logger
}

func (s *sender) Init(rt pkg.IsRuntime) {
	s.logger.Info().Msgf("emitting event in 3 seconds...")

	time.Sleep(time.Second * 3)

	rt.Events().Emit("test", "foo", 1, 2.3, []int{4, 5}).Wait()
}

func (s *sender) Name() string {
	return "sender"
}
