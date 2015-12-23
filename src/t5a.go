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

		str = inputString
		// ab, cd, pq, xy

		bb := str[0:2]
		ba := str[0]
		for i := 1; i < len(str)-1; i++ {
			if str[i+1] == ba {
				b1++
				//fmt.Printf("ba: %s %x %x\n", inputString, ba, str[i+1])
			}
			ba = str[i]

			if strings.Count(str, string(bb)) > 1 {
				b2++
				//fmt.Printf("bb: %s %s\n", inputString, bb)
			}

			if string(bb[0]) == string(bb[1]) {
				if strings.Count(str, string(bb)) > 1 && strings.Count(str, string(bb)) < 3 {
					b2--
					fmt.Printf("b2: %s %x %x\n", inputString, bb[0], bb[1])
				}
			}

			if i < len(str)-2 {
				bb = str[i : i+2]
			}

		}
		if b2 > 0 && b1 > 0 {
			s++
		}
	}
}
