# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Structure

Each day's puzzle solutions are organized in their own directory (`day1/`, `day2/`, etc.). Both Part 1 and Part 2 solutions are implemented in the same `main.go` file within each day's folder.

## Testing

Unit tests are included for edge cases and more challenging puzzles. You can find these in `main_test.go` files alongside the solutions.

## Philosophy

This repository focuses on solving puzzles and learning. The code is refactored when needed, but production-ready code isn't the primary goal here. The emphasis is on problem-solving and exploring different approaches.

## Running Solutions

```bash
# Run a specific day
cd day1
go run main.go

# Run tests
go test ./...
```
