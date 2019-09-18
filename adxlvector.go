package gpio

import "math"

// Gravity in [m/s^2]
const gravity = 9.80665

// The typical scale factor in mG/LSB for the 2g setting
// https://morf.lv/mems-part-1-guide-to-using-accelerometer-adxl345
const scaleFactor = 0.0039

// Precision of the float
const precision = 5

// Vector contains all the axes that the Adxl345 measures on
// It's raw data, so use the Get<method> functions to get usefull data
type Vector struct {
	x int16
	y int16
	z int16
}

func NewVector(block []byte) *Vector {
	a := new(Vector)

	//  raw.x = (((int)_buff[1]) << 8) | _buff[0];
	//  raw.y = (((int)_buff[3]) << 8) | _buff[2];
	//  raw.z = (((int)_buff[5]) << 8) | _buff[4];

	// x = 0x32 0x33
	a.x = (int16(block[1]) << 8) | int16(block[0])
	// y = 0x34 0x35
	a.y = (int16(block[3]) << 8) | int16(block[2])
	// z = 0x36 0x37
	a.z = (int16(block[5]) << 8) | int16(block[4])
	return a
}

// returns
// (1) the G Force on the x axe
// (2) the G Force on the y axe
// (3) the G Force on the z axe.
// If the Adxl345 is standing still in a flat position the z value ≈ 9.81
func (v *Vector) GetGforce() (float64, float64, float64) {
	x := round(float64(v.x) * scaleFactor)
	y := round(float64(v.y) * scaleFactor)
	z := round(float64(v.z) * scaleFactor)

	return x, y, z
}

func (v *Vector) GetMs2() (float64, float64, float64) {
	xgf, ygf, zgf := v.GetGforce()

	x := round(xgf * gravity)
	y := round(ygf * gravity)
	z := round(zgf * gravity)

	return x, y, z
}

// Pitch and Roll actually depend on the way the accelerometer is placed.
// This function is written based on the normal position of a prototype board
func (v *Vector) GetPitch() float64 {
	x, y, z := v.GetGforce()
	pitch := (math.Atan2(y, math.Sqrt(x*x+z*z)) * 180.0) / math.Pi
	return pitch
}

// Pitch and Roll actually depend on the way the accelerometer is placed.
// This function is written based on the normal position of a prototype board
func (v *Vector) GetRoll() float64 {
	x, y, z := v.GetGforce()
	roll := (math.Atan2(x, math.Sqrt(y*y+z*z)) * 180.0) / math.Pi
	return roll
}

func round(f float64) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(f*shift+.5) / shift
}
