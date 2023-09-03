package main

import (
	"github.com/gravestench/runtime"
)

// the listener service will react to instances of this being added

type exampleService struct {
	name string
}

func (e *exampleService) Init(rt runtime.Runtime) {
	// noop
}

func (e *exampleService) Name() string {
	return e.name
}
