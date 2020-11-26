package main

import (
	"io"
	"math/rand"
	"os"
	"time"
)

func openFileAppendly(filename string) (io.Writer, func()) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return f, func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}
}

func randomNumBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

type void struct{}

func randomTimerBetween(a, b int) chan void {
	ch := make(chan void)
	go func() {
		var loop func()
		loop = func() {
			n := randomNumBetween(a, b)
			time.Sleep(time.Duration(n) * time.Second)
			ch <- void{}
			loop()
		}
		loop()
	}()
	return ch
}
