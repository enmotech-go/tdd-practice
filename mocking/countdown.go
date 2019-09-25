package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigSleeper struct {
	Duration time.Duration
}

func (c ConfigSleeper) Sleep() {
	time.Sleep(c.Duration)
}

const finalword = "go!"
const countstart = 3

func Countdown(b io.Writer, sleeper Sleeper) {
	for i := countstart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(b, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(b, finalword)
}

func main() {
	sleeper := &ConfigSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
