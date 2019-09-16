package gpio

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Led struct {
	pin rpio.Pin
}

// Param pinNum corresponds to the BCM Pin Number
// Returns a pointer to a Led
func NewLed(pinNum uint8) *Led {
	l := new(Led)
	l.pin = rpio.Pin(pinNum)
	// Set pin to output mode
	l.pin.Output()
	return l
}

func (l *Led) On() {
	l.pin.High()
}

func (l *Led) Off() {
	l.pin.Low()
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
