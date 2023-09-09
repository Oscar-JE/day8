package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParsePositiveNumbers(fileName string) [][]int {
	readFile, _ := os.Open(fileName)
	file_scanner := bufio.NewScanner(readFile)
	m := [][]int{}
	for file_scanner.Scan() {
		var inputLine string = file_scanner.Text()
		var line []int = []int{}
		for _, char := range inputLine {
			line = append(line, int(char)-48)
		}
		m = append(m, line)
	}
	return m
}

func convertToInt(number string) int {
	converted, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error during conversion")
		return -1
	}
	return converted
}

func convertToIntArr(numbers []string) []int {
	line := []int{}
	for _, el := range numbers {
		integer := convertToInt(el)
		line = append(line, integer)
	}
	return line
}
