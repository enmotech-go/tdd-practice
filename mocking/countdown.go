package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
	"time"
)

/**
class_09
demand: 从3每隔一秒钟递减1，递减到0输出GO!
*/

const (
	finalWorld     = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

func main() {
	sleeper := &ConfigurableSleeper{time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep()  {
// 	s.Calls++ // 监视器是一种Mock, 可以记录依赖关系怎么被调用
// }

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
