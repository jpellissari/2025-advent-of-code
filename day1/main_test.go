package main

import (
	"fmt"
	"testing"
)

func TestCrossings(t *testing.T) {
	testCases := []struct {
		name                 string
		movements            []movement
		expected             int
		expectedDialPosition int
	}{
		{
			name: "R movement: 1 cross and final position 10",
			movements: []movement{
				{"R", 60},
			},
			expected:             1,
			expectedDialPosition: 10,
		},
		{
			name: "L movement: 1 cross and final position at 90",
			movements: []movement{
				{"L", 60},
			},
			expected:             1,
			expectedDialPosition: 90,
		},
		{
			name: "count 2 crossings from multiple movements",
			movements: []movement{
				{"L", 68},
				{"L", 30},
				{"R", 48},
				{"L", 5},
			},
			expected:             2,
			expectedDialPosition: 95,
		},
		{
			name: "Count 10 crossing from one big movement R 1000",
			movements: []movement{
				{"R", 1000},
			},
			expected:             10,
			expectedDialPosition: 50,
		},
		{
			name: "Count 6 crossings from AoC example",
			movements: []movement{
				{"L", 68},
				{"L", 30},
				{"R", 48},
				{"L", 5},
				{"R", 60},
				{"L", 55},
				{"L", 1},
				{"L", 99},
				{"R", 14},
				{"L", 82},
			},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Running test case:", tc.name)
			crossings, finalPosition := executePart2(tc.movements)

			if crossings != tc.expected {
				fmt.Printf("Expected crossings: %d, but got: %d\n", tc.expected, crossings)
			}

			if tc.expectedDialPosition != 0 && finalPosition != tc.expectedDialPosition {
				fmt.Printf("Expected final dial position: %d, but got: %d\n", tc.expectedDialPosition, finalPosition)
			}
		})
	}
}
