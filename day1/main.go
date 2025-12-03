package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type movement struct {
	direction string
	distance  int
}

func main() {
	movements, err := parseInput()
	if err != nil {
		fmt.Printf("failed to parse input: %v", err)
		os.Exit(1)
	}

	password := executePart1(movements)
	crossings, _ := executePart2(movements)

	fmt.Println("Part1 result: ", password)
	fmt.Println("Part2 result: ", crossings)

}

func executePart1(m []movement) int {
	currentPosition := 50

	var password int
	for _, m := range m {
		if m.direction == "R" {
			currentPosition = (currentPosition + m.distance) % 100
		} else {
			currentPosition = (currentPosition - m.distance + 100) % 100
		}

		if currentPosition == 0 {
			password++
		}
	}

	return password
}

func executePart2(m []movement) (int, int) {
	currentPosition := 50
	crossings := 0

	for _, move := range m {
		var step int

		if move.direction == "R" {
			step = 1
		} else {
			step = -1
		}

		for i := 0; i < move.distance; i++ {
			currentPosition = (currentPosition + step + 100) % 100
			if currentPosition == 0 {
				crossings++
			}
		}
	}

	return crossings, currentPosition
}

func parseInput() ([]movement, error) {
	file, err := os.Open("./day1/input")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var movements []movement
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance in line %q: %w", line, err)
		}
		movements = append(movements, movement{direction, distance})
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return movements, nil
}
