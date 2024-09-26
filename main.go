package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("This program aims to give examples and a deeper understanding, for me at least, about algorithms and data structures\n\n")
	ArrayAddresses()
	report()
	memoryAllocation()
}

func StringMemoryAnalisys() {
	fmt.Println("\nNow strings:")
	var a string
	var b string
	fmt.Println("string a address: ", &a)
	fmt.Println("string b address: ", &b)
	a = "Babylon rule dem sdluhaslfdbsdflksdbfliasduhflsdakjbfnsdkjn"
	fmt.Println("string a new address: ", &a)
	c := "udhflksdjfbnsdalkjfhsdauifhsadljkfnbasdlkjfnsduipfhsdaijfbnaskdjfnbasdkjlfbnsadlkvjncblkxcjzvbniuadshfeuihfdksjnfskdmbnvsivujhbnslkjnflkdjsbniusdhgvdiufnbvldkfjbvndfiouvhpiouhsdpigfjbncxkvmnbpuiwshefp9u8hdsçkjvcnxzlkbnvccccowisuhbgfçldsjkohfsdkljfbn"
	var d string

	fmt.Println("string c address: ", &c)
	fmt.Println("string d address: ", &d)
	fmt.Println("\n O que eu observei disso: Memoria é alocada sequencialmente conforme o desenvolvimento da funcao e um string geralmente ocupa 10 bytes de memoria")
}

func ArrayAddresses() {
	fmt.Println("Chapter 1: Arrays\n")
	fmt.Println("Arrays are a contigous memory space that is alocated by multiplying the size of its basic type by the number of items")
	var a [5]int
	fmt.Println(a)
	arrayAddress := &a
	i0Address := &a[0]
	i1Address := &a[1]

	fmt.Printf("endereço do array: %p\n", arrayAddress)
	fmt.Println("endereço do primeiro item do array: ", i0Address)
	fmt.Println("endereço do proximo item do array: ", i1Address)
	fmt.Println("tipo de um pointer para um int:", reflect.TypeOf(i0Address))

	addressString0 := fmt.Sprintf("%v", i0Address)
	addressString1 := fmt.Sprintf("%v", i1Address)

	intAddress0, err := convertHexStringToInt(addressString0)
	if err != nil {
		log.Fatal(err)
	}
	intAddress1, err := convertHexStringToInt(addressString1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tamanho de um int no array: ", intAddress1-intAddress0)
}

func ArrayMemoryAnalisys() {
	fmt.Println("Lets start with arrays:")
	var a [5]int
	fmt.Println("array:", a)
	fmt.Println("What happens when an array item value extrapolates the size defined by its type?")
	//Memory overflow
	a[2] = int(math.Pow(2, 64))
	firstHexString := fmt.Sprintf("%v", &a[2])
	firstAddress, err := convertHexStringToInt(firstHexString)
	fmt.Println("endereço do numero grande: ", firstAddress)
	if err != nil {
		log.Fatal("Error converting string in memoryAllocation")
	}

	secondHexString := fmt.Sprintf("%v", &a[3])
	secondAddress, err := convertHexStringToInt(secondHexString)
	fmt.Println("endereço do proximo item do array: ", secondAddress)
	a2Size := secondAddress - firstAddress
	fmt.Println("array:", a)
	fmt.Println("a2Size: ", a2Size)

}

func memoryAllocation() {
	fmt.Println("\n\nChapter 2: Memory allocation\n")
	fmt.Println("Memory allocation defines how to allocate memory, wich types have wich size and how different data strcutures work around memory")
	fmt.Println("The objective of this chapter is to get a basic grasp of memory, but every chapter hereafter is getting its own memory section")

	ArrayMemoryAnalisys()
	StringMemoryAnalisys()
}

func convertHexStringToInt(s string) (int, error) {
	var sum float64
	for i := 2; i < len(s); i++ {
		char := fmt.Sprintf("%c", s[i])
		var number float64
		power := float64(len(s) - 1 - i)

		switch string(char) {
		case "a":
			number = 10 * (math.Pow(16, power))
		case "b":
			number = 11 * (math.Pow(16, power))
		case "c":
			number = 12 * (math.Pow(16, power))
		case "d":
			number = 13 * (math.Pow(16, power))
		case "e":
			number = 14 * (math.Pow(16, power))
		case "f":
			number = 15 * (math.Pow(16, power))
		default:
			convertedInt, err := strconv.Atoi(char)
			if err != nil {
				return -1, err
			}
			convertedFloat := float64(convertedInt)
			number = convertedFloat * (math.Pow(16, power))
		}
		sum = sum + number
	}
	return int(sum), nil
}
