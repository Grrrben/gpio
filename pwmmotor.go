/**
 * A work in progress, just going forward at the moment
 */
package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"math"
)

const frequency = 64000

// L293D is 4 channel driver that creates the logic to
// drive a Motor. The pwmMotor makes use of the pwm
// to be able to give more or less power to the motor
type PwmMotor struct {
	pinPlus   rpio.Pin
	pinMin    rpio.Pin
	pinEnable rpio.Pin

	// to current speed
	DutyLen  uint32
	cycleLen uint32
}

// NewMotor factory. Returns a pointer to a new PwmMotor instance
func NewPwmMotor(plus, min, enable int) *PwmMotor {
	m := new(PwmMotor)

	m.DutyLen = 0
	m.cycleLen = 32

	m.pinEnable = rpio.Pin(enable)
	m.pinEnable.Output()
	m.pinEnable.Low()

	m.pinPlus = rpio.Pin(plus)
	m.pinPlus.Output()
	m.pinPlus.Mode(rpio.Pwm)
	m.pinPlus.Freq(frequency)
	m.pinPlus.DutyCycle(m.DutyLen, m.cycleLen)
	m.pinPlus.High()

	m.pinMin = rpio.Pin(min)
	m.pinMin.Output()
	//m.pinMin.Mode(rpio.Pwm)
	//m.pinMin.Freq(frequency)
	//m.pinMin.DutyCycle(m.DutyLen, m.cycleLen)
	m.pinMin.Low()

	return m
}

func (m *PwmMotor) SpinClockwize() {
	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()
}

//func (m *PwmMotor) SpinCounterClockwize() {
//	m.pinPlus.Low()
//	m.pinMin.High()
//	m.pinEnable.High()
//
//}

// Stop disables the output to the motor and sets the duty cycle to 0
func (m *PwmMotor) Stop() {
	m.pinEnable.Low()
	m.pinPlus.DutyCycle(0, m.cycleLen)
	m.pinMin.DutyCycle(0, m.cycleLen)
}

// IsSpinning returns a bool that tell's if the motor is running.
func (m *PwmMotor) IsSpinning() bool {
	return rpio.ReadPin(m.pinEnable) == rpio.High
}

func (m *PwmMotor) Faster() {
	if m.DutyLen < 31 {
		m.DutyLen = m.DutyLen + 2
	}

	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()

	m.pinPlus.DutyCycle(m.DutyLen, m.cycleLen)
	//fmt.Println("Set pinPlus to ", m.DutyLen)
	//
	//
	//if rpio.ReadPin(m.pinPlus) == rpio.High {
	//	fmt.Println("plus is High")
	//}
	//if rpio.ReadPin(m.pinMin) == rpio.High {
	//	fmt.Println("min is High")
	//}
	//
	////if rpio.ReadPin(m.pinPlus) == rpio.High {
	//	m.pinPlus.DutyCycle(m.DutyLen, m.cycleLen)
	//	fmt.Println("Set pinPlus to ", m.DutyLen)
	//	return
	//}
	//m.pinMin.DutyCycle(m.DutyLen, m.cycleLen)
	//fmt.Println("Set pinMin to ", m.DutyLen)
}

func (m *PwmMotor) Slower() {
	if m.DutyLen > 0 {
		m.DutyLen = m.DutyLen - 2
	}

	if m.DutyLen < 0 {
		m.DutyLen = 0
	}

	m.pinPlus.DutyCycle(m.DutyLen, m.cycleLen)
}

func (m *PwmMotor) SetPwmPercentage(percentage uint32) {
	if percentage > 100 {
		percentage = 100
	}

	f := float64(m.cycleLen) / 100
	p := math.Round(f * float64(percentage))
	m.DutyLen = uint32(p)
	m.pinPlus.DutyCycle(m.DutyLen, m.cycleLen)
}
