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

	// initialize the HC-SR04 distance meter
	distancemeter := gpio.NewHCSR04(5, 6)
	for {
		cm := distancemeter.Measure()
		fmt.Printf("Distance in CM is %.2f\n", cm)
		time.Sleep(time.Second / 2)
	}
}
