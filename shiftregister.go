package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// a 74HC595 shiftregister
type ShiftRegister struct {
	dataPin rpio.Pin
	stcpPin rpio.Pin
	shcpPin rpio.Pin

	state rpio.State
	data  uint8
}

// NewShiftRegister is a factory for a new ShiftRegister instance
// the struct is based on the 74HC595 shiftregister hardware.
func NewShiftRegister(sdiNum, rclkNum, srclkNum int) (*ShiftRegister, error) {
	r := new(ShiftRegister)

	// serial data input
	r.dataPin = rpio.Pin(sdiNum)
	r.dataPin.Output()
	r.dataPin.Low()

	// memory clock input(STCP)
	// flush after data writes are done
	r.stcpPin = rpio.Pin(rclkNum)
	r.stcpPin.Output()
	r.stcpPin.Low()

	// shift register clock input(SHCP)
	// flush after each individual data write
	r.shcpPin = rpio.Pin(srclkNum)
	r.shcpPin.Output()
	r.shcpPin.Low()

	return r, nil
}

// flushShcp Flush the Shcp pin
// call after each individual data write
func (r *ShiftRegister) flushShcp() {
	r.shcpPin.Write(r.state ^ 0x01)
	r.shcpPin.Write(r.state)
}

// flushStcp Flush the Stcp pin
// call after data writes are done
func (r *ShiftRegister) flushStcp() {
	r.stcpPin.Write(r.state ^ 0x01)
	r.stcpPin.Write(r.state)
}

// setBit sets an individual bit
func (r *ShiftRegister) setBit(bit rpio.State) {
	r.dataPin.Write(bit)
	r.flushShcp()
}

// SendData sends the bytes to the shiftregister
func (r *ShiftRegister) SendData(data uint8) {
	r.data = data
	for i := uint(0); i < 8; i++ {
		r.setBit(rpio.State((r.data >> i) & 0x01))
	}
	r.flushStcp()
}
