package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var countdownStart = 3
var finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{
		duration: time.Second * 1,
		sleep:    time.Sleep,
	})
}
