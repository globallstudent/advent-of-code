package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Read input file and parse numbers into two slices
func readInput(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text()) // Split by spaces
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", scanner.Text())
		}

		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid numbers in line: %s", scanner.Text())
		}

		left = append(left, a)
		right = append(right, b)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

// Part 1: Calculate the sum of absolute differences
func minDifferenceSum(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff // Absolute value
		}
		sum += diff
	}

	return sum
}

// Part 2: Calculate similarity score
func similarityScore(left, right []int) int {
	// Create a frequency map for right list
	countRight := make(map[int]int)
	for _, num := range right {
		countRight[num]++
	}

	// Calculate similarity score
	score := 0
	for _, num := range left {
		score += num * countRight[num]
	}

	return score
}

func main() {
	filePath := "input.txt"

	// Read and parse input
	leftList, rightList, err := readInput(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Solve Part 1
	result1 := minDifferenceSum(leftList, rightList)
	fmt.Println(result1)

	// Solve Part 2
	result2 := similarityScore(leftList, rightList)
	fmt.Println(result2)
}
