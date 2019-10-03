package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

// a 74HC595 shiftregister
type ShiftRegister struct {
	sdi   rpio.Pin
	rclk  rpio.Pin
	srclk rpio.Pin
}

func NewShiftRegister(sdiNum, rclkNum, srclkNum int) (*ShiftRegister, error) {
	r := new(ShiftRegister)

	// serial data input
	r.sdi = rpio.Pin(sdiNum)
	r.sdi.Output()
	r.sdi.Low()

	// memory clock input(STCP)
	r.rclk = rpio.Pin(rclkNum)
	r.rclk.Output()
	r.rclk.Low()

	// shift register clock input(SHCP)
	r.srclk = rpio.Pin(srclkNum)
	r.srclk.Output()
	r.srclk.Low()

	return r, nil
}

// Shift a set of 8 bitsSet the data to 74HC595
func (r *ShiftRegister) Shift(b byte) {

	rpio.SpiTransmit(0x80 & b)

	r.srclk.High()
	time.Sleep(time.Millisecond * 50)
	r.srclk.Low()
	r.rclk.High()
	time.Sleep(time.Millisecond * 50)
	r.rclk.Low()

	//
	//rpio.SpiTransmit(b)
	//time.Sleep(time.Second)
	//rpio.SpiTransmit(0x80 & b)
	//time.Sleep(time.Second)

	//GPIO.output(SDI, 0x80 & (dat << bit))
	//GPIO.output(SRCLK, GPIO.HIGH)
	//time.sleep(0.001)
	//GPIO.output(SRCLK, GPIO.LOW)
	//GPIO.output(RCLK, GPIO.HIGH)
	//time.sleep(0.001)
	//GPIO.output(RCLK, GPIO.LOW)
}
