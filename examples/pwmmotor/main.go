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

	// initialize the motor
	// BCM Pin number 10 corresponds to physical pin number 19
	mtr := gpio.NewPwmMotor(19, 13, 26)

	// press ctrl-C to quit the program
	mtr.SpinClockwize()

	fmt.Println("start")

	for i := 0; i < 16; i++ {
		mtr.Faster()
		time.Sleep(time.Second / 2)
	}

	for j := 0; j < 16; j++ {
		mtr.Slower()
		time.Sleep(time.Second / 2)
	}

	time.Sleep(time.Second * 3)

	fmt.Println("starting percentages")

	mtr.SetPwmPercentage(25)
	time.Sleep(time.Second * 2)
	mtr.SetPwmPercentage(50)
	time.Sleep(time.Second * 2)
	mtr.SetPwmPercentage(75)
	time.Sleep(time.Second * 2)
	mtr.SetPwmPercentage(100)
	time.Sleep(time.Second * 2)

	mtr.Stop()
}
