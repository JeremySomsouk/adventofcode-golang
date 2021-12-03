package readerUtils

import (
	. "adventofcode-golang/Domain"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func getFile(day int) (*os.File, error) {
	filename := fmt.Sprintf("day%02d.txt", day)
	return os.Open(path.Join("resources", filename))
}

func ScanStrings(day int) []string {
	file, _ := getFile(day)
	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func ScanNumbers(day int) []int {
	file, _ := getFile(day)
	defer file.Close()

	var integers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, integer)
	}

	return integers
}

func ScanCommands(day int) []Command {
	file, _ := getFile(day)
	defer file.Close()

	var commands []Command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(command[1])

		commands = append(commands, Command{command[0], value})
	}

	return commands
}
