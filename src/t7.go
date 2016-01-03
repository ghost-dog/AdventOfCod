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

	var prog [340][6]string
	//var prog_t [5]string
	words := map[string]string{"RSHIFT": "RSHIFT", "OR": "OR", "AND": "AND", "NOT": "NOT", "LSHIFT": "LSHIFT", "->": "->"}

	type node struct {
		value int
		check bool
	}
	//var sign map[string]int
	sign := make(map[string]interface{})

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
	a := 0
	b := 0
	s := 0
	//j := 0
	exit := 0
	for exit == 0 {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			//fmt.Println("sign:", sign)
			//fmt.Printf("Total: %d %d", s, i)
			//return // error or EOF
			exit = 1
		}
		//fmt.Printf("The input was: %s", inputString)

		if _, err := fmt.Sscanf(inputString, "%s %s %s %s %s", &prog[i][0], &prog[i][1], &prog[i][2], &prog[i][3], &prog[i][4]); err == nil {
			//fmt.Printf("%s %s %s -> %s\n", prog[i][0], prog[i][1], prog[i][2], prog[i][3])
			if _, err := strconv.Atoi(prog[i][0]); err != nil {
				if _, found := sign[prog[i][0]]; !found {
					if _, found := words[prog[i][0]]; !found {
						sign[prog[i][0]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][1]); err != nil {
				if _, found := sign[prog[i][1]]; !found {
					if _, found := words[prog[i][1]]; !found {
						sign[prog[i][1]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][2]); err != nil {
				if _, found := sign[prog[i][2]]; !found {
					if _, found := words[prog[i][2]]; !found {
						sign[prog[i][2]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][3]); err != nil {
				if _, found := sign[prog[i][3]]; !found {
					if _, found := words[prog[i][3]]; !found {
						sign[prog[i][3]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][4]); err != nil {
				if _, found := sign[prog[i][4]]; !found {
					if _, found := words[prog[i][4]]; !found {
						sign[prog[i][4]] = node{0, false}
					}
				}
			}

			i++
		} else if _, err := fmt.Sscanf(inputString, "%s %s %s %s", &prog[i][0], &prog[i][1], &prog[i][2], &prog[i][3]); err == nil {
			if _, err := strconv.Atoi(prog[i][0]); err != nil {
				if _, found := sign[prog[i][0]]; !found {
					if _, found := words[prog[i][0]]; !found {
						sign[prog[i][0]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][1]); err != nil {
				if _, found := sign[prog[i][1]]; !found {
					if _, found := words[prog[i][1]]; !found {
						sign[prog[i][1]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][2]); err != nil {
				if _, found := sign[prog[i][2]]; !found {
					if _, found := words[prog[i][2]]; !found {
						sign[prog[i][2]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][3]); err != nil {
				if _, found := sign[prog[i][3]]; !found {
					if _, found := words[prog[i][3]]; !found {
						sign[prog[i][3]] = node{0, false}
					}
				}
			}
			i++
		} else if _, err := fmt.Sscanf(inputString, "%s %s %s", &prog[i][0], &prog[i][1], &prog[i][2]); err == nil {
			if _, err := strconv.Atoi(prog[i][0]); err != nil {
				if _, found := sign[prog[i][0]]; !found {
					if _, found := words[prog[i][0]]; !found {
						sign[prog[i][0]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][1]); err != nil {
				if _, found := sign[prog[i][1]]; !found {
					if _, found := words[prog[i][1]]; !found {
						sign[prog[i][1]] = node{0, false}
					}
				}
			}
			if _, err := strconv.Atoi(prog[i][2]); err != nil {
				if _, found := sign[prog[i][2]]; !found {
					if _, found := words[prog[i][2]]; !found {
						sign[prog[i][2]] = node{0, false}
					}
				}
			}
			i++
		}

	}
	// parsing
	exit = 0
	al := false
	bl := false
	sl := false
	for exit == 0 {
		for i := 0; i < 340; i++ {
			if prog[i][5] != "1" {
				a = 0
				al = false
				b = 0
				bl = false
				s = 0
				sl = false
				//fmt.Printf("%s %s %s %s %s\n", prog[i][0], prog[i][1], prog[i][2], prog[i][3], prog[i][4])
				if prog[i][1] == "->" {
					if n, err := strconv.Atoi(prog[i][0]); err == nil {
						a = n
						al = true
					} else {
						if nd, found := sign[prog[i][1]]; found {
							if nd.check {
								a = nd.value
								al = true
							}
						}
					}
					if al {
						if nd, found := sign[prog[i][2]]; found {
							nd.value = a
							nd.check = true
							prog[i][5] = "1"
						}
					}

				} else if prog[i][2] == "->" {

				} else if prog[i][3] == "->" {

				}
				if n, err := strconv.Atoi(prog[i][0]); err == nil {

				}

			}
		}
		if nd, found := sign[prog[i][1]]; found {
			if nd.check {
				fmt.Printf(" Line a %n\n", nd.value)
				exit = 1
			}
		}
	}
	//fmt.Println(sign)
}
