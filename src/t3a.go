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

	inputFile, inputError := os.Open("adv3.txt")
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
			return // error or EOF
		}
		//fmt.Printf("The input was: %s", inputString)
		var north int
		var south int
		var east  int
		var west  int
	
		str = inputString
		north = strings.Count(str, "^")
		south = strings.Count(str, "v")
		east  = strings.Count(str, ">")
		west  = strings.Count(str, "<")
		
		width := east+west
		height := north+south
		fmt.Printf("north : %d\n", north)
		fmt.Printf("south : %d\n", south)
		fmt.Printf("east  : %d\n", east)
		fmt.Printf("west  : %d\n", west)
		fmt.Printf("width : %d\n", width)
		fmt.Printf("height: %d\n", height)
		
	}
}



