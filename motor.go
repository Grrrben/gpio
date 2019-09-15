package gpio

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
)

// L293D is 4 channel driver that creates the logic to
// drive the motor

type motor struct {
	pinPlus   rpio.Pin
	pinMin    rpio.Pin
	pinEnable rpio.Pin
}

func NewMotor(plus, min, enable int) *motor {
	m := new(motor)
	m.pinPlus = rpio.Pin(plus)
	m.pinPlus.Output()

	m.pinMin = rpio.Pin(min)
	m.pinMin.Output()

	m.pinEnable = rpio.Pin(enable)
	m.pinEnable.Output()
	m.pinEnable.Low()

	return m
}

func (m *motor) Forwards() {
	fmt.Println("Starting to spin forwards")
	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()
}

func (m *motor) Backwards() {
	fmt.Println("Starting to spin backwards")
	m.pinPlus.Low()
	m.pinMin.High()
	m.pinEnable.High()
}

func (m *motor) Stop() {
	fmt.Println("Stoppping to spin")
	m.pinEnable.Low()
}
