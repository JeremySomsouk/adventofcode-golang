package readerUtils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func GetFile(day int) (*os.File, error) {
	filename := fmt.Sprintf("day%02d.txt", day)
	return os.Open(path.Join("resources", filename))
}

func ScanStrings(file *os.File) []string {
	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func ScanNumbers(file *os.File) []int {
	var integers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, integer)
	}

	return integers
}