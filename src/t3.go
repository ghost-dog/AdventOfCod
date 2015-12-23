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
		var path  [4120][4074]int

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				path[x][y] = 0
			}
		}
	
		l := west
		h := south
		path[l][h] = 1	
		for i := 0; i < len(str); i++ {
			if string(str[i]) == "^" {
				h++
			}
			if string(str[i]) == "v" {
				h--
			}
			if string(str[i]) == ">" {
				l++
			}
			if string(str[i]) == "<" {
				l--
			}
                        path[l][h] = path[l][h] + 1
			
		}
		total := 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if path[x][y] > 0 {
					total++
				}	
			}
		}
		fmt.Printf("total: %d\n", total)
		
	}
}



