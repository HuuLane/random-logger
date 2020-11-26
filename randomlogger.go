package main

import (
	"io"
	"log"
)

type RandomLogger struct {
	i *log.Logger
	w *log.Logger
	e *log.Logger
}

func NewRandomLogger(out io.Writer) *RandomLogger {
	rl := new(RandomLogger)
	rl.i = log.New(out, "Info: ", log.LstdFlags)
	rl.w = log.New(out, "Warning: ", log.LstdFlags)
	rl.e = log.New(out, "Error: ", log.LstdFlags)
	return rl
}

func (rl *RandomLogger) Info(v ...interface{}) {
	rl.i.Println(v...)
}
func (rl *RandomLogger) Warning(v ...interface{}) {
	rl.w.Println(v...)
}
func (rl *RandomLogger) Error(v ...interface{}) {
	rl.e.Println(v...)
}

func (rl *RandomLogger) RandomLog() {
	// possibility
	// 0.9  info
	// 0.09 warning
	// 0.01 error

	content := FakeData()

	n := randomNumBetween(1, 100)
	if n <= 90 {
		rl.Info(content)
		return
	}
	if n == 100 {
		rl.Error(content)
		return
	}
	rl.Warning(content)
}

func (rl *RandomLogger) Burst() {
	n := randomNumBetween(50, 200)
	for i := 0; i < n; i++ {
		rl.RandomLog()
	}
}
