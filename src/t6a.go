package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//	"strings"
)

func Even(number int) bool {
	return number%2 == 0
}

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	var str, str1, str2, str3, str4 string
	var a, b, a1, b1 int
	var light [1000][1000]int

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			light[x][y] = 0
		}
	}

	inputFile, inputError := os.Open("adv6.txt")
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
			total := 0
			for y := 0; y < 1000; y++ {
				for x := 0; x < 1000; x++ {
					total = total + light[x][y]
				}
			}
			fmt.Printf("total: %d\n", total)

			return // error or EOF
		}
		if _, err := fmt.Sscanf(inputString, "%s %s", &str, &str1); err == nil {
			//fmt.Printf("str:%s %s\n", str, str1)
			fmt.Printf("%s", inputString)
			if str == "toggle" {
				if _, err := fmt.Sscanf(inputString, "%s %3d,%3d %s %3d,%3d", &str2, &a, &b, &str3, &a1, &b1); err == nil {
					fmt.Printf("toggle : %d %d %d %d\n", a, b, a1, b1)
					for y := b; y <= b1; y++ {
						for x := a; x <= a1; x++ {
							light[x][y] = light[x][y] + 2
						}
					}

				}

			}
			if str == "turn" && str1 == "off" {
				if _, err := fmt.Sscanf(inputString, "%s %s %3d,%3d %s %3d,%3d", &str2, &str4, &a, &b, &str3, &a1, &b1); err == nil {
					fmt.Printf("turn of : %d %d %d %d\n", a, b, a1, b1)
					for y := b; y <= b1; y++ {
						for x := a; x <= a1; x++ {
							light[x][y]--
							if light[x][y] < 0 {
								light[x][y] = 0
							}
						}
					}

				}

			}

			if str == "turn" && str1 == "on" {
				if _, err := fmt.Sscanf(inputString, "%s %s %3d,%3d %s %3d,%3d", &str2, &str4, &a, &b, &str3, &a1, &b1); err == nil {
					fmt.Printf("turn on : %d %d %d %d\n", a, b, a1, b1)
					for y := b; y <= b1; y++ {
						for x := a; x <= a1; x++ {
							light[x][y]++
						}
					}

				}

			}
			//fmt.Printf("%dx%dx%d\n",l,w,h) // Outputs 123
		}

	}

}
