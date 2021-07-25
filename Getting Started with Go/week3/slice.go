package main

import (
	"fmt"
	"sort"
)

func main() {
	var ele rune
	var size int
	var sli = make([]int, 0, 3)
	size = cap(sli)
	for i := 0; i <= size; i++ {
		if i >= len(sli) {
			size = size + 1
		}
		ele = 0
		fmt.Println("Enter a number to add or x to exit: ")
		fmt.Scan(&ele)
		if ele == 0 {
			fmt.Println("Stopping!")
			break
		}
		sli = append(sli, int(ele))
		sort.Ints(sli)
		fmt.Println(sli)
	}

}
