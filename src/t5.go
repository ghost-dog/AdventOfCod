package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	var str string

	inputFile, inputError := os.Open("adv5.txt")
	s := 0
	i := 1
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Printf("Total: %d %d", s, i)
			return // error or EOF
		}
		//fmt.Printf("The input was: %s", inputString)
		b1 := 0
		b2 := 0
		b3 := 0

		str = inputString
		// ab, cd, pq, xy
		b3 = strings.Count(str, "ab") + strings.Count(str, "cd") + strings.Count(str, "pq") + strings.Count(str, "xy")
		b2 = strings.Count(str, "a") + strings.Count(str, "o") + strings.Count(str, "i") + strings.Count(str, "e") + strings.Count(str, "u")
		b1 = 0

		bb := str[0]
		for i := 1; i < len(str); i++ {
			if str[i] == bb {
				b1++
			}
			bb = str[i]
		}
		if b3 == 0 && b2 >= 3 && b1 > 0 {
			s++
		}
	}
}
