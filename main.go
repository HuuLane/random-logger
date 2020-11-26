package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	// Where?
	// check file flag
	// if flag, output to file
	// else to stdout
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

	// When?
	// burst:  per 120s-360s
	// normal: per 1s-20s
	normalTimer := randomTimerBetween(1, 20)
	burstTimer := randomTimerBetween(120, 360)
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
