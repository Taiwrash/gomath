package gomath

import (
	"log"
)

func Max(array []int) int {
	if len(array) <= 0 {
		log.Fatal("You passed an empty list")
		return -1
	}
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	return max
}

func Min(array []int) int {
	if len(array) <= 0 {
		log.Fatal("You passed an empty list")
		return -1
	}
	min := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}
