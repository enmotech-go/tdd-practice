package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	duration time.Duration
}

func (s *SpySleeper) Sleep() {
	time.Sleep(s.duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)

	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	//Countdown(os.Stdout,)
}
