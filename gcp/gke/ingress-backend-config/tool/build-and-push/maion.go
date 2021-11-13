package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("start")
}

func splitCmd(cmd string) []string {
	return strings.Split(cmd, " ")
}
