package gpio

import "github.com/corrupt/go-smbus"

const adxl345Address = 0x53
const adxl345RegDatax0 = 0x32   // X-axis data 0 (6 bytes for X/Y/Z)
const adxl345RegPowerCtl = 0x2D // Power-saving features control
const bandwidthRateAddress = 0x2c

const rangeAddress = 0x31
const adxl345Range2G = 0x00 // +/-  2g (default)

// Adxl345 is an accelerometer
// Some indepth knowledge about this component can be found on this web page:
// https://morf.lv/mems-part-1-guide-to-using-accelerometer-adxl345
type Adxl345 struct {
	bus         *smbus.SMBus
	Address     byte
	InterfaceId uint
}

// Returns a pointer to the Axdl345 and/or an error
func NewAdxl345(deviceId uint) (*Adxl345, error) {
	smb, err := smbus.New(deviceId, adxl345Address)

	adxl := new(Adxl345)
	adxl.bus = smb
	adxl.Address = adxl345Address
	adxl.InterfaceId = deviceId

	if err != nil {
		return adxl, err
	}

	if err = adxl.setDataRate(0x0A); err != nil {
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
func (a Adxl345) setDataRate(rate uint8) error {
	return a.bus.Write_byte_data(bandwidthRateAddress, rate)
}

// SetRange sets the (2g) range and checks for errors
func (a Adxl345) setRange() error {
	busData, err := a.bus.Read_byte_data(rangeAddress)
	if err != nil {
		return err
	}

	// thank you @ https://github.com/Devligue/go-adxl345/blob/master/go-adxl345.go
	value := int32(busData)
	value &= ^0x0F
	value |= int32(adxl345Range2G)
	value |= 0x08

	return a.bus.Write_byte_data(rangeAddress, byte(value))
}

// starts the measurement.
// Should always be called in the Factory
func (a Adxl345) start() error {
	return a.bus.Write_byte_data(adxl345RegPowerCtl, 0x08)
}

// GetVector reads the current data
// Returns a vector containing the raw XYZ data
func (a Adxl345) GetVector() (*Vector, error) {
	block := make([]byte, 6)
	_, err := a.bus.Read_i2c_block_data(adxl345RegDatax0, block)
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
