package main

import (
	"runtime"
)

func initPprof() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}
