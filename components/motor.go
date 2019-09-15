package components

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

// L293D is 4 channel driver that creates the logic to
// drive the Motor

type Motor struct {
	pinPlus   rpio.Pin
	pinMin    rpio.Pin
	pinEnable rpio.Pin
	locked    bool
}

func NewMotor(plus, min, enable int) *Motor {
	m := new(Motor)

	m.pinEnable = rpio.Pin(enable)
	m.pinEnable.Output()
	m.pinEnable.Low()

	m.pinPlus = rpio.Pin(plus)
	m.pinPlus.Output()
	m.pinPlus.High()

	m.pinMin = rpio.Pin(min)
	m.pinMin.Output()
	m.pinMin.Low()

	m.locked = false

	return m
}

func (m *Motor) SpinClockwize() {
	fmt.Println("Starting to spin CW")
	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()
}

func (m *Motor) SpinCountClockwize() {
	fmt.Println("Starting to spin CCW")
	m.pinPlus.Low()
	m.pinMin.High()
	m.pinEnable.High()
}

func (m *Motor) Stop() {
	fmt.Println("Stoppping to spin")
	m.pinEnable.Low()
}

func (m *Motor) IsSpinning() bool {
	if rpio.ReadPin(m.pinEnable) == rpio.High {
		return true
	}
	return false
}

func (m *Motor) Forwards() {
	if rpio.ReadPin(m.pinPlus) == rpio.High {
		// already going fw?
		if !m.IsSpinning() {
			m.pinEnable.High()
		}
	} else {
		// it's going backwards... Toggle the pins
		m.toggle()
	}
}

func (m *Motor) Backwards() {
	if rpio.ReadPin(m.pinMin) == rpio.High {
		// already going bw?
		if !m.IsSpinning() {
			m.pinEnable.High()
		}
	} else {
		// it's going forwards... Toggle the pins
		m.toggle()
	}
}

// From spinning CW to CCW to CW etc
func (m *Motor) toggle() {

	if m.locked {
		// already switching i presume?
		return
	}

	if m.IsSpinning() {
		m.Stop()
	}

	m.locked = true

	time.Sleep(time.Second / 2)

	fmt.Println("Switching")
	m.pinPlus.Toggle()
	m.pinMin.Toggle()
	m.pinEnable.High()

	m.locked = false
}
