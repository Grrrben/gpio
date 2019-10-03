package main

import (
	"github.com/grrrben/gpio"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"time"
)

func main() {

	if err := rpio.Open(); err != nil {
		panic(err)
	}

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0)

	register, err := gpio.NewShiftRegister(9, 10, 11)
	if err != nil {
		log.Fatal(err)
	}

	//bitset := []byte{0x01, 0x02, 0x04, 0x02, 0x01, 0x02, 0x04, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02, 0x01, 0x02}
	//bitset := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}

	//for i := 0; i < len(bitset); i++ {
	//	register.Shift(bitset[i])
	//
	//	fmt.Printf("byte %d transmitted\n", bitset[i])
	//	time.Sleep(time.Millisecond * 100)
	//}
	for {
		register.Shift(0x02)
		time.Sleep(time.Millisecond * 250)
		register.Shift(0x02)
		time.Sleep(time.Millisecond * 250)
		register.Shift(0x01)
		time.Sleep(time.Millisecond * 250)
		register.Shift(0x01)
		time.Sleep(time.Millisecond * 250)
		register.Shift(0x80)
		time.Sleep(time.Millisecond * 250)
		register.Shift(0x80)
		time.Sleep(time.Millisecond * 250)
	}

}
