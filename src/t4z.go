package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x\n", md5.Sum(data))
	data = []byte("abcdef609043")
	fmt.Printf("%x\n", md5.Sum(data))
	data = []byte("pqrstuv1048970")
	fmt.Printf("%x\n", md5.Sum(data))

}
