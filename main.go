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
	// burst:  per 1s-2min
	// normal: per 10min-25min
	normalTimer := randomTimer(1, 60*2)
	burstTimer := randomTimer(60*10, 60*25)
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
