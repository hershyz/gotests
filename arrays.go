package main

import (
	"fmt"
)

func main() {

	var arr = [5]int{0, 1, 2, 3, 4}
	var i int = 0
	for i < len(arr) {
		fmt.Println(arr[i])
		i++
	}
}
