package main

import (
	"adventofcode-golang/lib"
	"fmt"
)

func main() {
	file, _ := readerUtils.GetFile(1)
	defer file.Close()
	text := readerUtils.ScanNumbers(file)

	for _, value := range text {
		fmt.Println(value)
	}
}
