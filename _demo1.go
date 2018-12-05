package main

import (
	"github.com/j7b/gopherino/core"
	"github.com/j7b/gopherino/pin"
)

func main() {
	println("starting")
	pin.D13.Output()
	v := true
	for {
		pin.D13.Set(v)
		println("millis", core.Millis(), "micros", core.Micros(), "rand", core.Random())
		v = !v
		core.Delay(1000)
	}
}
