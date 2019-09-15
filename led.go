package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type led struct {
	pin rpio.Pin
}

// var pinNum corresponds to the BCM Pin Number
func NewLed(pinNum uint8) *led {
	l := new(led)
	l.pin = rpio.Pin(pinNum)
	// Set pin to output mode
	l.pin.Output()
	return l
}

func (l *led) Blink() {
	l.pin.High()
	time.Sleep(time.Second / 5)
	l.pin.Low()
}

// Blink a number of times
func (l *led) BlinkBlink(num int) {
	for x := 0; x < num; x++ {
		l.Blink()
		time.Sleep(time.Second / 5)
	}
}
