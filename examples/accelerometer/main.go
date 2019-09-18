package main

import (
	"fmt"
	"github.com/grrrben/gpio"
	"log"
	"time"
)

func main() {
	adxl, err := gpio.NewAdxl345(1)
	if err != nil {
		log.Fatal(err)
	}
	defer adxl.Close()

	for {
		// always get a fresh Vector to work with
		v, err := adxl.GetVector()
		if err != nil {
			log.Println("Failed to get the Vector of the adxl345;", err)
		}

		// print the g force and gravity
		GforceMs2(v)

		// print pitch and roll
		PitchAndRoll(v)

	}
}

func GforceMs2(v *gpio.Vector) {
	gx, gy, gz := v.GetGforce()
	mx, my, mz := v.GetMs2()

	fmt.Printf("G force xyz:   %.2f, %.2f, %.2f\n", gx, gy, gz)
	fmt.Printf("Ms2 xyz:                            %.2f, %.2f, %.2f\n", mx, my, mz)
	time.Sleep(time.Second / 5)
}
func PitchAndRoll(v *gpio.Vector) {
	fmt.Printf("Pitch %.2f and roll %.2f\n", v.GetPitch(), v.GetRoll())
	time.Sleep(time.Second / 5)
}
