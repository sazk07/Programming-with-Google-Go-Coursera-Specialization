package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"math"
)

func main() {
	fmt.Println("enter a floating point number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	floatvar, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	integervar := int(floatvar)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(integervar)
}
