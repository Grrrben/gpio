package gpio

import (
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
	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()
}

func (m *Motor) SpinCounterClockwize() {
	m.pinPlus.Low()
	m.pinMin.High()
	m.pinEnable.High()
}

func (m *Motor) Stop() {
	m.pinEnable.Low()
}

func (m *Motor) IsSpinning() bool {
	if rpio.ReadPin(m.pinEnable) == rpio.High {
		return true
	}
	return false
}

func (m *Motor) Clockwize() {
	if rpio.ReadPin(m.pinPlus) == rpio.High {
		// already going CW?
		if !m.IsSpinning() {
			m.pinEnable.High()
		}
	} else {
		// it's going CCW... Toggle the pins
		m.toggle()
	}
}

func (m *Motor) CounterClockwize() {
	if rpio.ReadPin(m.pinMin) == rpio.High {
		// already going CCW?
		if !m.IsSpinning() {
			m.pinEnable.High()
		}
	} else {
		// it's going CW... Toggle the pins
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

	// wait till the current switching stops...
	time.Sleep(time.Second / 2)

	m.pinPlus.Toggle()
	m.pinMin.Toggle()
	m.pinEnable.High()

	m.locked = false
}
