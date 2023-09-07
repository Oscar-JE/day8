package parser

import (
	"bufio"
	"os"
)

func ParseWood(fileName string) {
	readFile, _ := os.Open(fileName)
	file_scanner := bufio.NewScanner(readFile)
	for file_scanner.Scan() {
		var line string = file_scanner.Text()
		println(line)
	}

}
