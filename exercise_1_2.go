package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	var s, sep string
	sep = " "

	for i:=0; i< len(os.Args); i++{
		s = strconv.Itoa(i) + sep + os.Args[i]
		fmt.Println(s)
	}
}