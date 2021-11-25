package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/pkg/term"
)

func main() {
	term, err := term.Open("/dev/cu.usbmodem1422203", term.Speed(115200))
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

			fmt.Println(string(b))
		}
	}()

	pixel := 0
	for {
		if pixel > 11 {
			pixel = 0
		}
		data, err := json.Marshal(map[string]interface{}{
			"cmd":        "setlight",
			"pixel":      pixel,
			"brightness": 100,
			"color":      "green",
		})
		if err != nil {
			panic(err)
		}
		pixel++
		_, err = term.Write(data)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
