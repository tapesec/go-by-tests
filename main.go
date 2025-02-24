package main

import (
	"os"
	"time"

	"github.com/tapesec/go-by-tests/mock"
)

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {

	mock.Countdown(os.Stdout, DefaultSleeper{})
}
