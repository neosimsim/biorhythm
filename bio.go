package main

import (
	"math"
)

func scaledSin(x, interval float64) float64 {
	x = 2 * math.Pi / interval * x
	return math.Sin(x)
}

func biorythm(width, daysAlive int) (plots [3]*Plot) {
	phisical := make([]int, width)
	for i := 0; i < width; i++ {
		phisical[i] = int(10 * scaledSin(float64(daysAlive+i), 23))
	}

	emotional := make([]int, width)
	for i := 0; i < width; i++ {
		emotional[i] = int(10 * scaledSin(float64(daysAlive+i), 28))
	}

	intel := make([]int, width)
	for i := 0; i < width; i++ {
		intel[i] = int(10 * scaledSin(float64(daysAlive+i), 33))
	}

	plot := &Plot{values: phisical, c: 'P'}
	plot2 := &Plot{values: emotional, c: 'E'}
	plot3 := &Plot{values: intel, c: 'I'}
	return [...]*Plot{plot, plot2, plot3}
}
