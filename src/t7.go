package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// AND OR NOT LSHIFT RSHIFT

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	var prog [340][5]string
	type node struct {
		value int
		check bool
	}

	//var sign map[string]int
	sign := make(map[string]node)

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
	j := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			//fmt.Println("sign:", sign)
			fmt.Printf("Total: %d %d", s, i)
			return // error or EOF
		}
		//fmt.Printf("The input was: %s", inputString)

		if _, err := fmt.Sscanf(inputString, "%s %s %s -> %s", &prog[i][0], &prog[i][1], &prog[i][2], &prog[i][3]); err == nil {
			//fmt.Printf("%s %s %s -> %s\n", prog[i][0], prog[i][1], prog[i][2], prog[i][3])
			if n, err := strconv.Atoi(prog[i][0]); err == nil {
				sign[prog[i][0]].value = n
				sign[prog[i][0]].check = true
			} else {
				if _, found := sign[prog[i][0]]; !found {
					sign[prog[i][0]].value = 0
					sign[prog[i][0]].check = false
				}
			}
			if n, err := strconv.Atoi(prog[i][2]); err == nil {
				sign[prog[i][2]].value = n
				sign[prog[i][2]].check = true
			} else {
				if _, found := sign[prog[i][2]]; !found {
					sign[prog[i][2]].value = 0
					sign[prog[i][2]].check = false
				}
			}
			if _, found := sign[prog[i][3]]; !found {
				sign[prog[i][3]].value = 0
				sign[prog[i][3]].check = false
			}

			i++
		} else if _, err := fmt.Sscanf(inputString, "%s %s -> %s", &prog[i][0], &prog[i][1], &prog[i][3]); err == nil {
			if prog[i][0] == "NOT" {
				if _, found := sign[prog[i][1]]; !found {
					sign[prog[i][1]].value = 0
					sign[prog[i][1]].check = false
				}
				if _, found := sign[prog[i][3]]; !found {
					sign[prog[i][3]].value = 0
					sign[prog[i][3]].check = false
				}
			} //fmt.Printf("%s %s -> %s\n", prog[i][0], prog[i][1], prog[i][3])
			i++
		} else if _, err := fmt.Sscanf(inputString, "%s -> %s", &prog[i][0], &prog[i][3]); err == nil {
			//fmt.Printf("%s -> %s\n", prog[i][0], prog[i][3])

			if _, found := sign[prog[i][1]]; !found {
				sign[prog[i][1]].value = 0
				sign[prog[i][1]].check = false
			}
			if _, found := sign[prog[i][3]]; !found {
				sign[prog[i][3]].value = 0
				sign[prog[i][3]].check = false
			}
			i++
		}

	}
	// parsing
	//for k := 1;
}
