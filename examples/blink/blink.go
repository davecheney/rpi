package main

import (
	"github.com/davecheney/rpi"
	"time"
)

func main() {
	for {
		rpi.GPIO.Set(rpi.GPIO_P1_22)
		time.Sleep(100 * time.Millisecond)
		rpi.GPIO.Clear(rpi.GPIO_P1_22)
		time.Sleep(100 * time.Millisecond)
	}
}

