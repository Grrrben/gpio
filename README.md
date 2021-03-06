# Go ports for electronic components

This is a collection of golang bindings for electronic components that are useful for robotics.

[![GoDoc](https://godoc.org/github.com/Grrrben/gpio?status.svg)](https://godoc.org/github.com/Grrrben/gpio)

## Led

Blingbling. 

### Usage
See `examples/led/main.go` for usage.

## Button

Check if a button is pressed. 

### Usage
See `examples/button/main.go` for usage.

## Motor / L293D

The motor is driven by a L293D is motor controller chip that 
gives the motor the commands to move forward or backward.  

### Usage
Use the 3 public methods to spin CW, CWW or stop:
```
SpinClockwize()
SpinCounterClockwize()
Stop()
```

## HC-SR04 Distance meter

The HC-SR04 is a pulse/echo based distance meter. 
Utilising the `Measure()` method, it will return a `float64` representing a distance in CM's.

The circuit for connecting the HC-SR04 to the pi is a bit tricky to build. 
Just [search](https://duckduckgo.com/?q=HC-SR04+circuit+raspberry&t=ffab&iax=images&ia=images&iai=https%3A%2F%2Ftutorials-raspberrypi.de%2Fwp-content%2Fuploads%2F2014%2F05%2Fultraschall_Steckplatine.png) for a nice diagram online.

As reading the distance is done by a sound based sensor, the meter is a bit slow. 
Moreover, the code uses `time.sleep()` to make sure that the sensor is ready for use. 

Putting this sensor in a goroutine can be a good way to go. 
You can just keep polling the distance while not having your entire program idle 
through the process.

However, multiple sensors can interfere with each each other due to receiving each others echo's.

> If you cycle through the sensors rather than trying to use them simultaneously 
then there will be no problem with incorrectly received echoes. 
They have a range of about 2m, so 4m (there and back) at 330m/s is 12ms. 
Allow double this for good measure, so check one sensor every 25ms.

See [raspberrypi.org](https://www.raspberrypi.org/forums/viewtopic.php?t=188700).

## ADXL345 Accelerometer

An accelerometer is a device that measures the acceleration in a specific direction from gravity and movement.
The ADXL345 is a 3 axis accelerometer, basically it can measure acceleration in 3 directions simultaneously.

The code gives you methods to retreive G force, the Ms2 gravitational force and degrees of tilt and roll.

The speed/velocity/distance is a work in progress. 

### Usage

To use the measurements of the ADXL345 use the helper type Vector (`adxlvector.go`) which 
has methods to retrieve usefull data such as the G force or a pitch and roll.
 
The pitch and roll actually depend on the way the accelerometer is placed.
This function is written based on the normal position of a prototype board

See `examples/accelerometer/main.go`.

### Wiring

Label | Description | Usage
--- | --- | ---
| 5V | Supply voltage 5V | unused |  
| 3.3V | 3.3v | wired to 3.3v (pin 1) |
| GND | Ground | wired to Ground (pin 6) |  
| SCL | Serial Communications Clock | wired to SCL (BCM 3) |   
| SDA | Serial Data (I2C) | wired to SDA (BCM 2) |   
| CS | Chip Select | wired to 3.3v |  
| SDO | Ground | wired to Ground |  
| INT1 | Interrupt 1 | unused |  
| INT2 | Interrupt 2 | unused |

### Datasheet

See the [datasheet](https://www.analog.com/media/en/technical-documentation/data-sheets/ADXL345.pdf) for 
additional information and possible settings of the ADXL345.

## 74HC595 ShiftRegister

Go bindings for the 74HC595, a shiftregister that controls 8 pins.

### usage

See `examples/shiftregister/main.go`.

## Pin numbers

See [pinout.xyz](https://pinout.xyz/) for a handy overview of the pin numbering of the Raspberry Pi.
![map](https://pinout.xyz/resources/raspberry-pi-pinout.png)

## Building the components

The ADXL345 is using the Go bindings for the SM bus from the 
github.com/corrupt/go-smbus package which uses CGO.  
For cross compiling, the CGO_ENABLED and CC args should be set in the build command.  

For debian/linux it's something like: 
`env GOOS=linux GOARCH=arm GOARM=5 CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 $(GOBUILD) -o $(BINARY_NAME) -v`

The **gcc-arm-linux-gnueabi** is needed; the package is fount in apt (debian...): 
`sudo apt install gcc-arm-linux-gnueabi`


See the `makefile` in the `/examples/<component>` directories.