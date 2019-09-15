package main

import (
	"fmt"
	"github.com/grrrben/glog"
	"github.com/grrrben/gpio"
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

	gpio.Blink()

}
