package main

import "fmt"

// Voce esta num predio e quer saber qual é o menor andar para soltar
// a bola e ela quebrar
func CrystalBall() {
	array := []bool{true}
	fmt.Println(FindFirstTrue(array))
}

func FindFirstTrue(array []bool) int {
	if len(array) == 0 {
		return -1
	}
	if len(array) == 1 && array[0] {
		return 0
	}

	left := 0
	right := len(array)

	for left < right {
		var arrow int = (left + right) / 2
		if array[arrow] {
			//walk backwards until you find a false and then return i + 1
			return WalkBackwards(array, arrow)
		} else {
			left = arrow + 1
		}
	}
	return -1
}

func WalkBackwards(array []bool, i int) int {
	for i >= 0 {
		if i == 0 && array[i] {
			return 0
		}
		if !array[i] {
			return i + 1
		}
		i--
	}
	return -1
}
