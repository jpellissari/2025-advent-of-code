package main

import "testing"

func TestInvalidIdPart1(t *testing.T) {
	testeCases := []struct {
		name        string
		r           IdRange
		expectedSum int
	}{
		{
			name: "Invalid IDs between 11-22 are 11 and 22, sum 33",
			r: IdRange{
				Start: 11, End: 22,
			},
			expectedSum: 33,
		},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sumInvalidIdInRangePart1(tc.r)
			if result != tc.expectedSum {
				t.Errorf("expected %v, got %v", tc.expectedSum, result)
			}
		})
	}
}

func TestCheckRanges(t *testing.T) {
	testeCases := []struct {
		name        string
		r           []IdRange
		expectedSum int
	}{
		{
			name: "Check AoC example range",
			r: []IdRange{
				{Start: 11, End: 22},
				{Start: 95, End: 115},
				{Start: 998, End: 1012},
				{Start: 1188511880, End: 1188511890},
				{Start: 222220, End: 222224},
				{Start: 1698522, End: 1698528},
				{Start: 446443, End: 446449},
				{Start: 38593856, End: 38593862},
			},
			expectedSum: 1227775554,
		},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executePart1(tc.r)
			if result != tc.expectedSum {
				t.Errorf("expected %v, got %v", tc.expectedSum, result)
			}
		})
	}
}

func TestSumInvalidIdPart2(t *testing.T) {
	testeCases := []struct {
		name        string
		r           []IdRange
		expectedSum int
	}{
		{
			name: "Ignore cases where ID are all the same digit, but have multiple substrings",
			r: []IdRange{
				{Start: 2222, End: 2222},
			},
			expectedSum: 2222,
		},

		{
			name: "Check AoC example range for part2",
			r: []IdRange{
				{Start: 11, End: 22},
				{Start: 95, End: 115},
				{Start: 998, End: 1012},
				{Start: 1188511880, End: 1188511890},
				{Start: 222220, End: 222224},
				{Start: 1698522, End: 1698528},
				{Start: 446443, End: 446449},
				{Start: 38593856, End: 38593862},
				{Start: 565653, End: 565659},
				{Start: 824824821, End: 824824827},
				{Start: 2121212118, End: 2121212124},
			},
			expectedSum: 4174379265,
		},
	}

	for _, tc := range testeCases {
		t.Run(tc.name, func(t *testing.T) {
			result := executePart2(tc.r)
			if result != tc.expectedSum {
				t.Errorf("expected %v, got %v", tc.expectedSum, result)
			}
		})
	}
}
