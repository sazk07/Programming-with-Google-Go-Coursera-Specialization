package main

import (
	"fmt"
	"strconv"
)

func swap(x int, y int) (int, int) {
	x = x + y
	y = x - y
	x = x - y
	return x, y
}
func BubbleSort(sli []int) {
	var i int
	var j int
	for i = 0; i < len(sli)-1; i++ {
		for j = i; j < len(sli); j++ {
			if sli[j] < sli[i] {
				sli[i], sli[j] = swap(sli[i], sli[j])
			}
		}
	}
}
func main() {
	fmt.Println("Enter upto integers, enter 'x' or 'X' when done:")
	s := make([]int, 0)
	var i int
	var temp string
	for i = 0; i < 10; i++ {
		fmt.Scan(&temp)
		if temp == "x" || temp == "X" {
			break
		}
		x, _ := strconv.Atoi(temp)
		s = append(s, x)
	}
	BubbleSort(s)
	fmt.Println("Sorted array slice:")
	for i = 0; i < len(s); i++ {
		fmt.Print(strconv.Itoa(s[i]) + " ")
	}
	fmt.Print("\n")
}
