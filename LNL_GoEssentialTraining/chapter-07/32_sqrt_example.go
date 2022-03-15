package main

import (
	"errors"
)

// Common errors
var (
	ErrNegSqrt    = errors.New("sqrt of negative number")
	ErrNoSolution = errors.New("no solution found")
)
func Abs(value float64) float64{

	if value < 0 {
		return -value
	}
	return value
}

// Sqrt return the square root of the number
func Sqrt(value float64) (float64, error) {
	if value < 0.0 {
		return 0.0, ErrNegSqrt
	}

	if value == 0.0 {
		return 0.0, nil
	}

	guess, epsilon := 1.0, 0.00001
	for i := 0; i < 10000; i++ {
		if Abs(guess * guess - value) <= epsilon {
			return guess, nil
		}
		guess = (value / guess + guess) / 2.0
	}

	return 0.0, ErrNoSolution
}