package main

import (
	readerUtils "adventofcode-golang/lib"
	"fmt"
	"strconv"
)

func part1(strings []string) int64 {
	var gammaBinary, epsilonBinary string

	bitSize := len(strings[0])
	for i := 0; i < bitSize; i++ {
		var counter int
		for _, element := range strings {
			switch string(element[i]) {
			case "1":
				counter++
			case "0":
				counter--
			}
		}
		if counter >= 0 {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	return gammaRate * epsilonRate
}

func part2(strings []string) int {
	return 0
}

func main() {
	strings := readerUtils.ScanStrings(3)

	fmt.Println("Result part 1 is :", part1(strings))
	fmt.Println("Result part 2 is :", part2(strings))
}
