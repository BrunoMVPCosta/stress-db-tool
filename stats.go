package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

var durations []int

func RegisterStats(ch chan int, chPrintStats chan bool) {

	go printMetrics()

	for {
		select {
		case duration := <-ch:
			durations = append(durations, duration)
		case <-chPrintStats:
			printMetrics()
			durations = nil
		}
	}
}

func printMetrics() {
	if len(durations) == 0 {
		fmt.Println("Nothing to report")
		return
	}
	sort.Ints(durations)

	index := int(Round(0.99*float64(len(durations)), .1, 0))

	aux := 0
	for i := index; i <= len(durations); i++ {
		aux += durations[i-1]
	}
	aux1 := aux / (len(durations) - index + 1)
	nineth := time.Duration(aux1) * time.Nanosecond

	sum := 0
	for i := range durations {
		sum += durations[i]
	}
	avg := sum / len(durations)

	fmt.Println("99th last 1 second ", nineth)
	fmt.Println("Average time ", time.Duration(avg)*time.Nanosecond)
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
