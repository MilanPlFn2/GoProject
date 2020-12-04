package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(text[0])

	for i := 0; i < len(text); i++ {
		text2 := strings.Split(text[i], " ")
		fmt.Println(text2[len(text2)-1])
	}
}
