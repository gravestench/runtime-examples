package main

import (
	"time"

	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	go func() {
		time.Sleep(time.Second * 5)
		rt.Shutdown()
	}()

	rt.Add(&listensForNewServices{})
	time.Sleep(time.Second)
	rt.Add(&exampleService{name: "foo"})
	time.Sleep(time.Second)
	rt.Add(&exampleService{name: "bar"})
	time.Sleep(time.Second)
	rt.Add(&exampleService{name: "baz"})

	rt.Run()
}
