package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func extractDigits(line string) int {
	var first, last rune
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			last = char
		}
	}
	if first == 0 || last == 0 {
		return 0
	}
	return int(first-'0')*10 + int(last-'0') // Convert to two-digit number
}

func calculateSum(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		totalSum += extractDigits(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return totalSum, nil
}

func main() {
	filePath := "input.txt"
	result, err := calculateSum(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
