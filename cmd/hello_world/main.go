package main

import "fmt"

func main() {
	printable := "Hello World"
	printed(printable)
}

func printed(printValue string) {
	fmt.Println(printValue)
}
