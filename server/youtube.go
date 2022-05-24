package server

import (
	"fmt"
	"time"
)

func RunMonitorTick( interval *time.Ticker ) {

	for {
		select {
		case <- interval.C:
			fmt.Println("ticker test")
		}
	}

}
