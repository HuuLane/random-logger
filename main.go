package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func openFileAppendly(filename string) (io.Writer, func()) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
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

func main() {
	// Where?
	// check file flag
	// if flag, output to file
	// else to stdout

	// When?
	// burst: 100-200 records per 50s-120s
	// normal: log a record per 1s-10s

	// How?
	// buffer
	// append mode
	// possibility
	// 0.9  info
	// 0.09 warning
	// 0.01 error

	out2file := flag.Bool("f", false, "output to file `random.log`")
	flag.Parse()

	var out io.Writer
	if *out2file {
		f, closer := openFileAppendly("random.log")
		defer closer()
		out = f
	} else {
		out = os.Stdout
	}

	L := NewRandomLogger(out)

	normalTimer := randomTimerBetween(1, 10)
	burstTimer := randomTimerBetween(50, 120)
	defer func() {
		close(normalTimer)
		close(burstTimer)
	}()
	for {
		select {
		case <-normalTimer:
			L.RandomLog()
		case <-burstTimer:
			L.Burst()
		}
	}
}
