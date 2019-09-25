package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
