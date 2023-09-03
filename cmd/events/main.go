package main

import (
	"time"

	"github.com/gravestench/runtime"
)

func main() {
	rt := runtime.New()

	rt.SetLogLevel(-1)

	rt.Add(&sender{})
	rt.Add(&receiver{})

	go func() {
		time.Sleep(time.Second * 5)
		rt.Shutdown()
	}()

	rt.Run()
}
