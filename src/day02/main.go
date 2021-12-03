package main

import (
	. "adventofcode-golang/Domain"
	readerUtils "adventofcode-golang/lib"
	"fmt"
)

func part1(commands []Command) int {
	var x, y int

	for _, command := range commands {
		moveValue := command.Value
		switch command.Name {
		case "forward":
			x += moveValue
		case "up":
			y -= moveValue
		case "down":
			y += moveValue
		}
	}

	return x * y
}

func part2(commands []Command) int {
	var x, y, aim int

	for _, command := range commands {
		moveValue := command.Value
		switch command.Name {
		case "forward":
			x += moveValue
			y += aim * moveValue
		case "up":
			aim -= moveValue
		case "down":
			aim += moveValue
		}
	}

	return x * y
}

func main() {
	commands := readerUtils.ScanCommands(2)

	fmt.Println("Result part 1 is :", part1(commands))
	fmt.Println("Result part 2 is :", part2(commands))
}
