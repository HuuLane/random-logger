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

	n := randomNumBetween(1, 100)
	if n <= 90 {
		rl.Info("This is a info")
		return
	}
	if n == 100 {
		rl.Error("This is a error")
		return
	}
	rl.Warning("This is a warning")
}

func (rl *RandomLogger) Burst() {
	n := randomNumBetween(50, 200)
	for i := 0; i < n; i++ {
		rl.RandomLog()
	}
}
