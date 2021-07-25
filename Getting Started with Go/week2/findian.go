package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	result := strings.Index(input, "i")
	if result == 0 {
		result2 := input[len(input)-1]
		if result2 == 'n' {
			result3 := strings.Contains(input, "a")
			if result3 == true {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}
}
