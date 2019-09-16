# Go ports for electronic components

This is a collection of different golang ports 
of several electronic components.

## Led

## Motor / L293D

The motor is driven by a L293D is 4 channel driver.



## HC-SR04 Distance meter

The HC-SR04 is a pulse/echo based distance meter.

## ADXL345 Accelerometer

The ADXL345 is using the Go bindings for the SM bus form the 
github.com/corrupt/go-smbus package which uses CGO.  
For cross compiling, the CGO_ENABLED and CC args should be set in the build command.  

For debian/linux it's something like: 
`env GOOS=linux GOARCH=arm GOARM=5 CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 $(GOBUILD) -o $(BINARY_NAME) -v`

See the `makefile` in the `/examples/accelerometer` directory.
Work in progress