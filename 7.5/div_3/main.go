package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if CheckDiv3(i) {
			fmt.Println(i)
		}
	}
}

func CheckDiv3(val int) bool {
	return val%3 == 0
}
