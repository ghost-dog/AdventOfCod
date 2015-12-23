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

	inputFile, inputError := os.Open("elves.txt")
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
			fmt.Printf("Total: %d %d", s,i)
			return // error or EOF
		}
		//fmt.Printf("The input was: %s", inputString)
		var l,w,h int
		if _, err := fmt.Sscanf(inputString, "%5dx%5dx%5d", &l, &w, &h ); err == nil {
			i = i + 1	
			s = s + 2*l*w + 2*l*h + 2*w*h

			if (l >= h) && (l >= w) {
				s = s + h*w
			} else if (w >= h) && (w >= l) {
				s = s + h*l
			} else if (h >= w) && (h >= l) {
				s = s + l*w
			}	

		    //fmt.Printf("%dx%dx%d\n",l,w,h) // Outputs 123
		}
	}
}
