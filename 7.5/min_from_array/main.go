package main

import (
	"fmt"
	"strconv"
)

func main() {
	var min int
	var min_pos int
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	min, min_pos = MinFromArray(x)

	fmt.Println("Minimal position: " + strconv.Itoa(min_pos))
	fmt.Println("Minimal value: " + strconv.Itoa(min))
}

func MinFromArray(arr []int) (min int, min_pos int) {
	for i, val := range arr {
		if (i == 0) || (val < min) {
			min = val
			min_pos = i
		}
	}
	return
}
