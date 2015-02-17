package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, stats!")
}

// Stream processing experiments
func MeanStream(in, out chan float64) {
	total := float64(0)
	count := float64(0)
	for {
		newn := <-in
		total += newn
		count++
		out <- (total / count)
	}
}

type PartialMean struct {
	total float64
	count int
}

func CalcPartialMean(xs []float64) PartialMean {
	total := float64(0)
	count := 0
	for _, f := range xs {
		total += f
		count++
	}
	return PartialMean{total, count}
}

// Basic statistical function on slices of float64
func mean(xs []float64) float64 {
	total := float64(0)
	for _, f := range xs {
		total += f
	}
	total /= float64(len(xs))
	return total
}

func variance(xs []float64) float64 {
	mean := mean(xs)
	diffs := float64(0)
	for _, f := range xs {
		diffs += math.Pow((mean - f), 2)
	}
	diffs /= float64(len(xs) - 1)
	return diffs
}

func sigma(xs []float64) float64 {
	return math.Sqrt(variance(xs))
}
