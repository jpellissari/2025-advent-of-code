package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	batteries, err := parseInput()
	if err != nil {
		fmt.Printf("failed to parse input: %v", err)
		os.Exit(1)
	}

	jolts := 0
	for _, bank := range batteries {
		jolts += getLargestJoltageFromBank(bank)

	}
	fmt.Println("Total jolts: ", jolts)
}

func getLargestJoltageFromBank(bank []int) int {
	// Find the largest digit that isn't in the last position
	largest := -1
	largestIdx := -1
	for i, digit := range bank {
		if i == len(bank)-1 {
			break
		}
		if digit > largest {
			largest = digit
			largestIdx = i
		}
	}

	// Find the second largest digit, starting AFTER the largest
	secondLargest := -1
	startIdx := largestIdx + 1

	for i := startIdx; i < len(bank); i++ {
		if bank[i] > secondLargest {
			secondLargest = bank[i]
		}
	}

	return largest*10 + secondLargest
}

func parseInput() ([][]int, error) {
	file, err := os.Open("./day3/input")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var batteries [][]int
	for scanner.Scan() {
		batteryJolts := scanner.Text()
		bank := make([]int, len(batteryJolts))
		for i := 0; i < len(batteryJolts); i++ {
			jolt, _ := strconv.Atoi(string(batteryJolts[i]))
			bank[i] = jolt
		}

		batteries = append(batteries, bank)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return batteries, nil
}
