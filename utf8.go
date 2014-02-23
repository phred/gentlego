package main

import (
	"fmt"
	"strings"
)

func δ(subj string) (string) {
	return strings.Replace(subj, "world", "💩", -1)
}

func main() {
//	dook := "💩"
//	ü := "dook"

	fmt.Println(δ("Hello, world!"))
}