package gpio

import (
	"fmt"
	"github.com/grrrben/gpio/components"
	"time"
)

type lighting struct {
	front *components.Led
	back  *components.Led
}

type distanceSensors struct {
	front *components.HCSR04
	back  *components.HCSR04
}

// a wrapper to control all different components of a car
type car struct {
	lights  lighting
	sensors distanceSensors
	// propulsion
	motor *components.Motor
}

func NewCar() *car {
	c := new(car)

	c.lights.front = components.NewLed(10)
	c.lights.back = components.NewLed(9)

	c.sensors.front = components.NewEcho(8, 7)
	c.sensors.back = components.NewEcho(20, 21)

	c.motor = components.NewMotor(17, 18, 27)

	return c
}

func (c *car) Init() {

	c.lights.front.Blink()
	c.lights.back.Blink()

	c.drive()

}

func (c *car) drive() {
	front := make(chan float64, 3)
	back := make(chan float64, 3)

	go func() {
		for {
			front <- c.sensors.front.Measure()
			time.Sleep(time.Second / 2)
		}

	}()

	go func() {
		for {
			back <- c.sensors.back.Measure()
			time.Sleep(time.Second / 2)
		}

	}()

	// only take action if both readings are known
	f := false
	b := false

	// caching last distances for taking measurements
	var lastFront, lastBack float64

	for {
		select {
		case cmFront := <-front:
			f = true
			lastFront = cmFront
			if f && b {
				diff := lastBack - lastFront
				if diff < 0 {
					c.forwards()
					fmt.Println("MOVE FORWARDS, ", lastBack, " cm in back")
				}
			}
		case cmBack := <-back:
			b = true
			lastBack = cmBack
			if f && b {
				diff := lastFront - lastBack
				if diff < 0 {
					c.backwards()
					fmt.Println("MOVE BACKWARDS, ", lastFront, " cm in front")
				}
			}
		}
	}
}

func (c *car) forwards() {
	c.motor.Forwards()
}

func (c *car) backwards() {
	c.motor.Backwards()
}
