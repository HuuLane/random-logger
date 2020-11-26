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
			log.Println("Send")
			ch <- void{}
			loop()
		}
		loop()
	}()
	return ch
}

func main() {
	// Where?
	// check file flag
	// if flag, output to file
	// else to stdout

	// When?
	// burst: 100-200 records per 30s-60s
	// normal: log a record per 1s-3s

	// How?
	// buffer
	// append mode

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

	I := log.New(out, "info: ", log.LstdFlags)
	W := log.New(out, "warning: ", log.LstdFlags)
	E := log.New(out, "error: ", log.LstdFlags)

	normalTimer := randomTimerBetween(1, 3)
	defer close(normalTimer)
	for {
		select {
		case <-normalTimer:
			I.Println("This is a info")
			W.Println("This is a warning")
			E.Println("This is a error")
		}
	}
}
