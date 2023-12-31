package main

import (
	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	// each service has a dependency that is not
	// actually resolved through the runtime but by
	// some other means (that part is up to you).
	rt.Add(newServiceWithAsyncDependencyResolution())
	rt.Add(newServiceWithAsyncDependencyResolution())
	rt.Add(newServiceWithAsyncDependencyResolution())
	rt.Add(newServiceWithAsyncDependencyResolution())
	rt.Add(newServiceWithAsyncDependencyResolution())

	rt.Run()
}
