package gpio

import "github.com/stianeikeland/go-rpio/v4"

type Button struct {
	pin rpio.Pin
}

func NewButton(pinNumber int) *Button {
	b := new(Button)
	b.pin = rpio.Pin(pinNumber)
	b.pin.Input()
	b.pin.PullUp()
	return b
}

// IsActive tell's whether the button is pressed.
func (b *Button) IsActive() bool {
	if b.pin.Read() == 0 {
		return true
	}
	return false
}
