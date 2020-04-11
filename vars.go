package main

import "fmt"

func main() {
	
	//Variable types: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64

	var num1 float64 = 1.5
	var num2 float64 = 2.5
	var run bool = true

	if run {
		fmt.Println(num1 + num2)
	}
}
