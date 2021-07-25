package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input numbers(separate with space):")
	newNumbers := bufio.NewReader(os.Stdin)
	lineRead, _, _ := newNumbers.ReadLine()
	splitNumStrings := strings.Split(string(lineRead), " ")
	var numValues []int
	for _, numStringsToConvert := range splitNumStrings {
		n, _ := strconv.Atoi(numStringsToConvert)
		numValues = append(numValues, n)
	}
	BubbleSort(numValues)
	fmt.Println(numValues)
}

func BubbleSort(enteredNumber []int) {
	for i := 0; i < len(enteredNumber); i++ {
		for j := 0; j < len(enteredNumber)-1-i; j++ {
			if enteredNumber[j+1] < enteredNumber[j] {
				Swap(enteredNumber, j)
			}
		}
	}
}

func Swap(comparisonNumber []int, j int) {
	var temp int
	temp = comparisonNumber[j]
	comparisonNumber[j] = comparisonNumber[j+1]
	comparisonNumber[j+1] = temp
}
