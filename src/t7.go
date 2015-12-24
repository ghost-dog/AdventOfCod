package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	var prog [340][5]string

	inputFile, inputError := os.Open("adv7.txt")

	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	i := 0
	s := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Printf("Total: %d %d", s, i)
			return // error or EOF
		}
		//fmt.Printf("The input was: %s", inputString)

		if _, err := fmt.Sscanf(inputString, "%s %s %s -> %s", &prog[i][0], &prog[i][1], &prog[i][2], &prog[i][3]); err == nil {
			fmt.Printf("%s %s %s -> %s\n", prog[i][0], prog[i][1], prog[i][2], prog[i][3])
			i++
		} else if _, err := fmt.Sscanf(inputString, "%s %s -> %s", &prog[i][0], &prog[i][1], &prog[i][3]); err == nil {
			fmt.Printf("%s %s -> %s\n", prog[i][0], prog[i][1], prog[i][3])
			i++
		} else if _, err := fmt.Sscanf(inputString, "%s -> %s", &prog[i][0], &prog[i][3]); err == nil {
			fmt.Printf("%s -> %s\n", prog[i][0], prog[i][3])
			i++
		}
	}
}
