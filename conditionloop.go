package main

import (
	"fmt"
)

func main() {

	var i float64 = 0
	
	for {
		fmt.Println(i)
		i = i + 0.1
		if i > 10.0 {
			break
		}
	}
}
