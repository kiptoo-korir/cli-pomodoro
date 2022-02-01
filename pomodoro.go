package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type remainingTime struct {
	total int
	min   int
	sec   int
}

func main() {
	fmt.Println("Pomodoro Starting Up...")

	var workingTime int
	var breakTime int
	var cycles int

	flag.IntVar(&workingTime, "work", 25, "Defines the amount of time in minutes that you would like to work.")
	flag.IntVar(&breakTime, "break", 5, "Defines the amount of time in minutes that you would to break from work.")
	flag.IntVar(&cycles, "cycles", 4, "Defines the number of pomodoro cycles you would like.")

	flag.Parse()

	fmt.Println("Working Time ", workingTime)
	fmt.Println("Resting Time", breakTime)
	fmt.Println("Cycles", cycles)

	for i := 0; i < cycles; i++ {
		startWork(workingTime)

		if (i + 1) == cycles {
			break
		}

		takeBreak(breakTime)
	}

	fmt.Println("\nEnd Of Cycle")
	os.Exit(0)
}

func startWork(amount int) {
	startTime := time.Now()
	endTime := startTime.Add(time.Minute * time.Duration(amount))

	for range time.Tick(1 * time.Second) {
		timeRemaining := getRemainingTime(endTime)

		fmt.Printf("\rWorking Minutes: %02d Seconds: %02d", timeRemaining.min, timeRemaining.sec)

		if timeRemaining.total <= 0 {
			break
		}
	}
}

func takeBreak(amount int) {
	breakStart := time.Now()
	breakEnd := breakStart.Add(time.Minute * time.Duration(amount))

	for range time.Tick(1 * time.Second) {
		breakRemaining := getRemainingTime(breakEnd)
		fmt.Printf("\rResting Minutes: %02d Seconds: %02d", breakRemaining.min, breakRemaining.sec)

		if breakRemaining.total <= 0 {
			break
		}
	}
}

func getRemainingTime(t time.Time) remainingTime {
	currentTime := time.Now()
	timeDiff := t.Sub(currentTime)

	totalSeconds := int(timeDiff.Seconds())
	minutes := int(totalSeconds/60) % 60
	seconds := int(totalSeconds % 60)

	return remainingTime{
		total: totalSeconds,
		min:   minutes,
		sec:   seconds,
	}
}
