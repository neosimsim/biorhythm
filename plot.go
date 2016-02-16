package main

import (
	"fmt"
)

type color string

const (
	DEFAULT color = "\x1b[0;0m"
	RED     color = "\x1b[0;31m"
	GREEN   color = "\x1b[0;32m"
	BLUE    color = "\x1b[0;34m"
	width         = 10
)

func printTimes(symb rune, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%c", symb)
	}
}

func plot(value float64) {
	discreteValue := int(value * width)
	if value > 0 {
		printTimes('.', width)
	} else {
		printTimes('.', width+discreteValue)
		printTimes('#', -discreteValue)
	}
	fmt.Printf("|")
	if value <= 0 {
		printTimes('.', width)
	} else {
		printTimes('#', discreteValue)
		printTimes('.', width-discreteValue)
	}
}

func plotValues(legend []string, values ...float64) {
	for index, value := range values {
		fmt.Print(legend[index], ":")
		plot(value)
		fmt.Print("  ")
	}
	fmt.Print("\n")
}
