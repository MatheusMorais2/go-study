package main

func linearSearch(array []int, number int) int {
	for i := 0; i < len(array); i++ {
		if number == array[i] {
			return i
		}
	}
	return -1
}
