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

	// initialize the btn
	// BCM pin number 18 corresponds to physical pin number 12
	btn := gpio.NewButton(18)

	isDown := false
	start := time.Now()

	// press ctrl-C to quit the program
	for {
		if btn.IsActive() {
			if !isDown {
				fmt.Println("button pressed")
				start = time.Now()
			}
			isDown = true
			continue
		}

		if isDown {
			elapsed := time.Since(start)
			fmt.Printf("button released after %d ms\n", elapsed.Milliseconds())
		}
		isDown = false
		time.Sleep(time.Second / 10)
	}
}
