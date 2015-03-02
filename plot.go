package main

import (
	"fmt"
	"strconv"
)

type Plot struct {
	values []int
	c      rune
}

func maxInt(values []int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func minInt(values []int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maxOverPlots(plots []*Plot) int {
	max := plots[0].values[0]
	for _, plot := range plots {
		value := maxInt(plot.values)
		if value > max {
			max = value
		}
	}
	return max
}

func minOverPlots(plots []*Plot) int {
	min := plots[0].values[0]
	for _, plot := range plots {
		value := minInt(plot.values)
		if value < min {
			min = value
		}
	}
	return min
}

func printRune(c rune) {
	quoted := strconv.QuoteRuneToASCII(c)
	unquoted := quoted[1 : len(quoted)-1]
	fmt.Printf(unquoted)
}

func firstPlotWithValueInLine(index, line int, plots []*Plot) (plot *Plot) {
	for _, plot := range plots {
		if plot.values[index] == line {
			return plot
		}
	}
	return nil
}

func printLineOfPlots(width, line int, plots []*Plot) {
	for i := 0; i < width; i++ {
		plot := firstPlotWithValueInLine(i, line, plots)
		if plot == nil {
			fmt.Print(" ")
		} else {
			printRune(plot.c)
		}
	}
}

func printPlot(width int, plots ...*Plot) {
	max := maxOverPlots(plots)
	min := minOverPlots(plots)
	for line := max; line >= min; line-- {
		printLineOfPlots(width, line, plots)
	}
}
