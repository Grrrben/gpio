package main

import (
	"fmt"
	"github.com/grrrben/gpio"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"time"
)

func main() {

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	// initialize the led
	// BCM Pin number 10 corresponds to physical pin number 19
	led := gpio.NewLed(10)

	// this will just keep blinking.
	// press ctrl-C to quit the program
	for {
		led.Blink()
		time.Sleep(time.Second / 2)
	}
}
