package main

import "fmt"

func main() {
	printable := "Hello World"
	printed(printable)
	fmt.Println(intDivider(10, 3))
}

func printed(printValue string) {
	fmt.Println(printValue)
}

func intDivider(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
