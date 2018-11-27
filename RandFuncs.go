package main

import (
	"math/rand"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randFloat64(min float64, max float64) float64 {
	return float64(int(min) + rand.Intn(int(max)-int(min)))
}
