# Go ports for electronic components

This is a collection of different golang ports 
of several electronic components.

## Led

Blingbling. See `examples/led/main.go` for usage.

## Motor / L293D

The motor is driven by a L293D is motor controller chip that 
gives the motor the commands to move forward or backward.  

Use the 3 public methods to spin CW, CWW or stop:
```
SpinClockwize
SpinCounterClockwize
Stop
```

## HC-SR04 Distance meter

The HC-SR04 is a pulse/echo based distance meter. 
Utilising the Measure() method, it will return a `float64` representing a distance in CM's.

## ADXL345 Accelerometer

An accelerometer is a device that measures the acceleration in a specific direction from gravity and movement.
The ADXL345 is a 3 axis accel, basically it can measure acceleration in 3 directions simultaneously.
When placed on a flat surface will always measure 9.81m/s2.

To use the measurements of the ADXL345 use the helper type Vector (`adxlvector.go`) which has methods to retrieve
usefull data such as the G force or a pitch and roll.
 
The pitch and roll actually depend on the way the accelerometer is placed.
This function is written based on the normal position of a prototype board

The speed/velocity/distance is a work in progress. 

See `examples/accelerometer/main.go` for usage.

## Pin numbers

See [pinout.xyz/(https://pinout.xyz/) for a handy overview of the pin numbering of the Raspberry Pi.
![map](https://pinout.xyz/resources/raspberry-pi-pinout.png)

## Building the components

The ADXL345 is using the Go bindings for the SM bus from the 
github.com/corrupt/go-smbus package which uses CGO.  
For cross compiling, the CGO_ENABLED and CC args should be set in the build command.  

For debian/linux it's something like: 
`env GOOS=linux GOARCH=arm GOARM=5 CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 $(GOBUILD) -o $(BINARY_NAME) -v`

See the `makefile` in the `/examples/<component>` directories.