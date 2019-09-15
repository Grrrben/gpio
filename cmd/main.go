package main

import (
	"fmt"
	"github.com/grrrben/glog"
	"github.com/grrrben/gpio"
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Could not set a logdir. Msg %s", err)
	}

	glog.SetLogFile(fmt.Sprintf("%s/../log/gpio.log", dir))
	glog.SetLogLevel(glog.Log_level_info)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	car := gpio.NewCar()
	car.Init()

	// wait for everything to finish
	forever := make(chan bool)
	<-forever
}
