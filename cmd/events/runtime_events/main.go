package main

import (
	"time"

	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	go func() {
		time.Sleep(time.Second)
		rt.Shutdown()
	}()

	rt.Add(&listensForNewServices{})
	rt.Add(&exampleService{name: "foo"})
	rt.Add(&exampleService{name: "bar"})
	rt.Add(&exampleService{name: "baz"})

	rt.Run()
}
