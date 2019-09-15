package components

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Led struct {
	pin rpio.Pin
}

// var pinNum corresponds to the BCM Pin Number
func NewLed(pinNum uint8) *Led {
	l := new(Led)
	l.pin = rpio.Pin(pinNum)
	// Set pin to output mode
	l.pin.Output()
	return l
}

func (l *Led) Blink() {
	l.pin.High()
	time.Sleep(time.Second / 5)
	l.pin.Low()
}

// Blink a number of times
func (l *Led) BlinkBlink(num int) {
	for x := 0; x < num; x++ {
		l.Blink()
		time.Sleep(time.Second / 5)
	}
}
