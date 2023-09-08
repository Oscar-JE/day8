package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParsePositiveNumbers(fileName string) [][]int {
	readFile, _ := os.Open(fileName)
	file_scanner := bufio.NewScanner(readFile)
	m := [][]int{}
	for file_scanner.Scan() {
		var line string = file_scanner.Text()
		splitted := strings.Split(line, " ")
		lineInt := convertToIntArr(splitted)
		m = append(m, lineInt)
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
