package gpio

import (
	"fmt"
	"time"
)

func Measure() {
	front := make(chan float64, 3)
	back := make(chan float64, 3)

	frontEcho := NewEcho(8, 7)
	backEcho := NewEcho(20, 21)

	go func() {
		for {
			front <- frontEcho.Measure()
			time.Sleep(time.Second / 2)
		}

	}()

	go func() {
		for {
			back <- backEcho.Measure()
			time.Sleep(time.Second / 2)
		}

	}()

	// only take action if both readings are known
	f := false
	b := false

	// caching last distances for taking measurements
	var lastFront, lastBack float64

	for {
		select {
		case cmFront := <-front:
			f = true
			lastFront = cmFront
			if f && b {
				diff := lastBack - lastFront
				if diff < 0 {
					fmt.Println("MOVE FORWARDS, ", lastBack, " cm in back")
				}
			}
		case cmBack := <-back:
			b = true
			lastBack = cmBack
			if f && b {
				diff := lastFront - lastBack
				if diff < 0 {
					fmt.Println("MOVE BACKWARDS, ", lastFront, " cm in front")
				}
			}
		}
	}
}
