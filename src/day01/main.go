package main

import (
	"adventofcode-golang/lib"
	"fmt"
)

func sumSlice(numbers []int) int {
	var counter int

	for i := range numbers {
		counter += numbers[i]
	}

	return counter
}

func part1(numbers []int) int {
	var counter int

	previousValue := numbers[0]
	for _, value := range numbers[1:] {
		if value > previousValue {
			counter++
		}
		previousValue = value
	}

	return counter
}

func part2(numbers []int) int {
	var counter int

	for i := 0; i <= len(numbers)-4; i++ {
		if len(numbers[i+1 : i+4]) != 3 {
			break
		}

		currentSlice := sumSlice(numbers[i : i+3])
		nextSlice := sumSlice(numbers[i+1 : i+4])
		if nextSlice > currentSlice {
			counter++
		}
	}

	return counter
}

func main() {
	numbers := readerUtils.ScanNumbers(1)

	fmt.Println("Result part 1 is :", part1(numbers))
	fmt.Println("Result part 2 is :", part2(numbers))
}
