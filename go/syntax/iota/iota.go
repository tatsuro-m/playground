package main

import "fmt"

const (
	first = iota
	second
	third
)

const (
	_ = iota
	resetFirst
)

func main() {
	fmt.Println("iota")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(resetFirst)
}
