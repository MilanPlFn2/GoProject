package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 1000; i++ {
		fmt.Println("D" + strconv.Itoa(i) + ": { D:0 }")
	}
}
