package main

import (
	"github.com/davecheney/rpi"
	"time"
)

func main() {
	// set GPIO25 to output mode
	rpi.GPIOFSel(rpi.GPIO_P1_22, rpi.BCM2835_GPIO_FSEL_OUTP)
	// turn the led off on exit
	defer rpi.GPIOClear(rpi.GPIO_P1_22)
	for {
		rpi.GPIOSet(rpi.GPIO_P1_22)
		time.Sleep(10 * time.Millisecond)
		rpi.GPIOClear(rpi.GPIO_P1_22)
		time.Sleep(10 * time.Millisecond)
	}
}
