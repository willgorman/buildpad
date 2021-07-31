package main

import "github.com/pkg/term"

func main() {
	term, err := term.Open("/dev/cu.usbmodem1412203", term.Speed(115200))
	if err != nil {
		panic(err)
	}

	_, err = term.Write([]byte{byte('a')})
	if err != nil {
		panic(err)
	}
}
