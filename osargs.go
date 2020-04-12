package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	var i int = 0
	for i < len(args) {
		fmt.Println(args[i])
		i++
	}
}
