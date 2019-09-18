/**
 * https://morf.lv/mems-part-1-guide-to-using-accelerometer-adxl345
 */
package gpio

import (
	"github.com/corrupt/go-smbus"
)

const ADXL345_ADDRESS = 0x53
const ADXL345_REG_DATAX0 = 0x32      // X-axis data 0 (6 bytes for X/Y/Z)
const ADXL345_REG_POWER_CTL = 0x2D   // Power-saving features control
const ADXL345_DATARATE_100_HZ = 0x2c // default 100hz bandwidth
const ADXL345_RANGE_2_G = 0x00       // +/-  2g (default)

// Adxl345 is an accelerometer
type Adxl345 struct {
	bus         *smbus.SMBus
	Address     byte
	InterfaceId uint
}

// Returns a pointer to the Axdl345 and/or an error
func NewAdxl345(deviceId uint) (*Adxl345, error) {
	smb, err := smbus.New(deviceId, ADXL345_ADDRESS)

	adxl := new(Adxl345)
	adxl.bus = smb
	adxl.Address = ADXL345_ADDRESS
	adxl.InterfaceId = deviceId

	if err != nil {
		return adxl, err
	}

	if err = adxl.setDataRate(); err != nil {
		return adxl, err
	}

	if err = adxl.setRange(); err != nil {
		return adxl, err
	}

	if err = adxl.start(); err != nil {
		return adxl, err
	}

	return adxl, err
}

// Setting the data rate and checks for errors
// See https://github.com/sunfounder/SunFounder_Super_Kit_V3.0_for_Raspberry_Pi/blob/master/Python/17_adxl345.py
//
// ADXL345_DATARATE_0_10_HZ = 0x00
// ADXL345_DATARATE_0_20_HZ = 0x01
// ADXL345_DATARATE_0_39_HZ = 0x02
// ADXL345_DATARATE_0_78_HZ = 0x03
// ADXL345_DATARATE_1_56_HZ = 0x04
// ADXL345_DATARATE_3_13_HZ = 0x05
// ADXL345_DATARATE_6_25HZ = 0x06
// ADXL345_DATARATE_12_5_HZ = 0x07
// ADXL345_DATARATE_25_HZ = 0x08
// ADXL345_DATARATE_50_HZ = 0x09
// ADXL345_DATARATE_100_HZ = 0x0A // (default)
// ADXL345_DATARATE_200_HZ = 0x0B
// ADXL345_DATARATE_400_HZ = 0x0C
// ADXL345_DATARATE_800_HZ = 0x0D
// ADXL345_DATARATE_1600_HZ = 0x0E
func (a Adxl345) setDataRate() error {
	return a.bus.Write_byte_data(ADXL345_DATARATE_100_HZ, 0x0A)
}

// SetRange sets the (2g) range and checks for errors
// See https://github.com/sunfounder/SunFounder_Super_Kit_V3.0_for_Raspberry_Pi/blob/master/Python/17_adxl345.py
//
// ADXL345_RANGE_2_G = 0x00  // +/-  2g (default)
// ADXL345_RANGE_4_G = 0x01  // +/-  4g
// ADXL345_RANGE_8_G = 0x02  // +/-  8g
// ADXL345_RANGE_16_G = 0x03 // +/- 16g
func (a Adxl345) setRange() error {
	busData, err := a.bus.Read_byte_data(0x31)
	if err != nil {
		return err
	}

	// thank you @ https://github.com/Devligue/go-adxl345/blob/master/go-adxl345.go
	value := int32(busData)
	value &= ^0x0F
	value |= int32(ADXL345_RANGE_2_G)
	value |= 0x08

	return a.bus.Write_byte_data(0x31, byte(value))
}

// starts the measurement.
// Should always be called in the Factory
func (a Adxl345) start() error {
	return a.bus.Write_byte_data(ADXL345_REG_POWER_CTL, 0x08)
}

// GetVector reads the current data
// Returns a vector containing the raw XYZ data
func (a Adxl345) GetVector() (*Vector, error) {
	block := make([]byte, 6)
	_, err := a.bus.Read_i2c_block_data(ADXL345_REG_DATAX0, block)
	if err != nil {
		return nil, err
	}

	vector := NewVector(block)

	return vector, err
}

// Close closes the bus
func (a Adxl345) Close() {
	a.bus.Bus_close()
}
