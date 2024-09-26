package main

import (
	"fmt"
	"math/rand"
	"time"
)

func report() {
	testCases := []int{10, 100, 1000, 10000, 1000000, 10000000} 
	for i := 0; i < len(testCases); i++{
		CallBinarySearch(testCases[i])
	}
}

func CallBinarySearch(len int) {
	fmt.Println("\nFucking binary search:")
	array, target := generateBinarySearch(len)
	fmt.Println("array size: ", len)
	fmt.Println("target:", target)
	index, steps, timeTaken := BinarySearch(array, target)
	fmt.Println("index: ", index)
	fmt.Println("steps needed: ", steps)
	fmt.Println("timeTaken: ", timeTaken)
}

func generateBinarySearch(arraySize int) ([]float32, float32) {
	if arraySize < 1 {
		return nil, 0
	}
	array := make([]float32, arraySize)
	array[0] = rand.Float32()
	for i := 1; i < arraySize; i++ {
		array[i] = array[i-1] + rand.Float32()
	}
	targetIndex := rand.Intn(arraySize)
	return array, array[targetIndex]
}

func BinarySearch(array []float32, target float32) (index int, steps int, timeTaken time.Duration) {
	startTime := time.Now()
	var endTime time.Duration
	lenA := len(array)
	if lenA == 0 {
		endTime = time.Since(startTime)
		return -1, 0, endTime
	}
	if lenA == 1 {
		endTime = time.Since(startTime)
		return 1, 1, endTime
	}

	var left int = 0
	var right int = lenA
	for i := 1; i < lenA; i++ {
		var arrow int = (left + right) / 2
		if array[arrow] == target {
		endTime = time.Since(startTime)
			return arrow, i, endTime
		}
		if array[arrow] > target {
			right = arrow
			continue
		}
		if array[arrow] < target {
			left = arrow
			continue
		}

	}

	return -1, lenA, endTime
}
