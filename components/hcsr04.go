package components

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

// For the HC-SR04 ultrasonic distance meter
type HCSR04 struct {
	triggerPin rpio.Pin
	echoPin    rpio.Pin
}

func NewEcho(triggerNum, echoNum int) *HCSR04 {
	e := new(HCSR04)
	e.triggerPin = rpio.Pin(triggerNum)
	e.echoPin = rpio.Pin(echoNum)

	e.triggerPin.Output()
	e.echoPin.Input()

	return e
}

// returns the distance in centimeters
func (e *HCSR04) Measure() float64 {

	e.triggerPin.Low()
	time.Sleep(time.Microsecond * 30)
	e.triggerPin.High()
	time.Sleep(time.Microsecond * 30)
	e.triggerPin.Low()
	time.Sleep(time.Microsecond * 30)

	// sometimes the HC-SR04 stalls, if so we just break on a set timeout
	// todo check where it stalls (in which loop)
loopHigh:
	for timeout := time.After(time.Second); ; {
		select {
		case <-timeout:
			break loopHigh
		default:
		}
		status := e.echoPin.Read()
		if status == rpio.High {
			break
		}
	}

	begin := time.Now()

loopLow:
	for timeout := time.After(time.Second); ; {
		select {
		case <-timeout:
			break loopLow
		default:
		}
		status := e.echoPin.Read()
		if status == rpio.Low {
			break
		}
	}

	end := time.Now()
	diff := end.Sub(begin)
	cm := (float64(diff.Nanoseconds()) * 17150) / 1000000000
	//fmt.Println("DIFF", cm)

	return cm
}
