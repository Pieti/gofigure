package main

import "math"

// Given two crystal balls that will break if dropped from high enoguh
// distance, determine the exact spot in which it will break in the
// most optimized way.

func TwoCrystalBalls(breaks []bool) int {
	jump := int(math.Sqrt(float64(len(breaks))))

	for i := jump; i < len(breaks); i += jump {
		if breaks[i] {
			for j := i - jump; j < i; j++ {
				if breaks[j] {
					return j
				}
			}
		}
	}
	return -1
}
