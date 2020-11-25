package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	// Where?
	// check file flag
	// if flag, output to file
	// else to stdout

	// When?
	// burst: 100-200 records per 30s-1min
	// normal: log a record per 1s-3s

	// How?
	// buffer
	// append mode

	out2file := flag.Bool("file", false, "output to file `random.log`")
	flag.Parse()

	var out io.Writer
	if *out2file {
		f, err := os.OpenFile("random.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}()
		out = f
	} else {
		out = os.Stdout
	}

	I := log.New(out, "info: ", log.LstdFlags)
	W := log.New(out, "warning: ", log.LstdFlags)
	E := log.New(out, "error: ", log.LstdFlags)

	I.Println("This is info")
	W.Println("This is warning")
	E.Println("This is error")
}
