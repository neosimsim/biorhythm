package main

import (
	"math"
)

const (
	physicalInterval    = 23
	emotionalInterval   = 28
	intelectualInterval = 33
)

func scaledSin(x, interval float64) float64 {
	x = 2 * math.Pi / interval * x
	return math.Sin(x)
}

func biorythm(daysAlive int) (physical float64, emotional float64, intel float64) {
	physical = scaledSin(float64(daysAlive), physicalInterval)
	emotional = scaledSin(float64(daysAlive), emotionalInterval)
	intel = scaledSin(float64(daysAlive), intelectualInterval)
	return
}
