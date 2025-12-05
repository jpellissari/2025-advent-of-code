package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	Start int
	End   int
}

func main() {
	ranges, err := parseInput()
	if err != nil {
		fmt.Printf("failed to parse input: %v", err)
		os.Exit(1)
	}

	sumPart1 := executePart1(ranges)
	sumPart2 := executePart2(ranges)

	fmt.Println("Part 1: ", sumPart1)
	fmt.Println("Part 2: ", sumPart2)
}

func executePart1(ranges []IdRange) int {
	totalSum := 0

	for _, r := range ranges {
		totalSum += sumInvalidIdInRangePart1(r)
	}

	return totalSum
}

func sumInvalidIdInRangePart1(r IdRange) int {
	var sum int

	for i := r.Start; i <= r.End; i++ {
		s := strconv.Itoa(i)
		isEven := len(s)%2 == 0
		if isEven {
			p1 := s[:len(s)/2]
			p2 := s[len(s)/2:]

			if p1 == p2 {
				sum += i
			}
		}
	}

	return sum
}

func executePart2(ranges []IdRange) int {
	totalSum := 0

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			s := strconv.Itoa(i)
			if isInvalidIdPart2(s) {
				totalSum += i
			}
		}
	}

	return totalSum
}

func isInvalidIdPart2(s string) bool {
	n := len(s)

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			chunk := s[:i]
			invalid := true
			for j := i; j < n; j += i {
				if s[j:j+i] != chunk {
					invalid = false
					break
				}
			}
			if invalid {
				return true
			}
		}

	}
	return false
}

func parseInput() ([]IdRange, error) {
	content, err := os.ReadFile("./day2/input")
	if err != nil {
		return nil, err
	}

	var ranges []IdRange

	content = []byte(strings.TrimSpace(string(content)))

	rawRanges := strings.Split(string(content), ",")
	for _, r := range rawRanges {
		bounds := strings.Split(r, "-")

		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("failed to convert start id to int: %v", err)
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			fmt.Printf("failed to convert end id to int: %v", err)
		}

		ranges = append(ranges, IdRange{Start: start, End: end})
	}

	return ranges, nil
}
