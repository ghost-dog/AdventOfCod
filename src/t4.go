package main

import (
	"crypto/md5"
	"strconv"
	"fmt"
	"strings"
	"encoding/hex"
)

func main() {
	key := "bgvyzdsv"
	var b [16]byte	
	for i := 1 ; i > 0; i++ {
//        	data = []byte(strings.Join([]string{key, strconv.Itoa(i)},""))
		b = md5.Sum([]byte(strings.Join([]string{key, strconv.Itoa(i)},"")))
                //fmt.Printf("%d  %x\n",i,b)
		//fmt.Printf("%s\n", hex.EncodeToString(b[0:3])[:5])
		if i%1000 == 0 {
			fmt.Printf("%d  %x\n",i,b)
		}
		if strings.Compare(hex.EncodeToString(b[0:3])[:6],"000000") == 0 {
	        	fmt.Printf("%x\n", md5.Sum([]byte(strings.Join([]string{key, strconv.Itoa(i)},""))))
			fmt.Printf("%d\n", i)
			break
		}
	}	

}
