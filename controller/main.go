package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/pkg/term"
)

func main() {
	term, err := term.Open("/dev/cu.usbmodem1412203", term.Speed(115200))
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			b := make([]byte, 6)
			_, err = term.Read(b)
			if err != nil {
				panic(err)
			}
			{
				currfunc, currfile, currline, _ := runtime.Caller(0)
				fmt.Printf("Func: %s, File: %s, Line: %d\n", runtime.FuncForPC(currfunc).Name(), currfile, currline)
			}

			fmt.Println(b)
		}
	}()

	for {

		_, err = term.Write([]byte("a\n"))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
