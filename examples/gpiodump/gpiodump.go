package main

import (
	"fmt"
	"github.com/davecheney/rpi"
)

func main() {
	var pin uint8
	for ; pin < 32 ; pin++ {
		fmt.Printf("GPIO%d: %v\n", pin, rpi.GPIO.Get(pin))
	}
}
