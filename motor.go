package gpio

import "github.com/stianeikeland/go-rpio"

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

func (m *motor) Spin() {
	m.pinPlus.High()
	m.pinMin.Low()
	m.pinEnable.High()
}

/**

GPIO.output(MotorPin1, GPIO.HIGH)
GPIO.output(MotorPin2, GPIO.LOW)
# Enable the motor
GPIO.output(MotorEnable, GPIO.HIGH)

import RPi.GPIO as GPIO


# Set up pins
MotorPin1 = 17
MotorPin2 = 18
MotorEnable = 27


def setup():
    # Set the GPIO modes to BCM Numbering
    GPIO.setmode(GPIO.BCM)
    # Set pins to output
    GPIO.setup(MotorPin1, GPIO.OUT)
    GPIO.setup(MotorPin2, GPIO.OUT)
    GPIO.setup(MotorEnable, GPIO.OUT, initial=GPIO.LOW)



*/
