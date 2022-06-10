package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	if str == "" {
		fmt.Printf("%d", 0)
		return
	}
	words := strings.Split(str, " ")
	fmt.Printf("%d", len(words))
}
