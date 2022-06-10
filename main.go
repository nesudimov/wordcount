package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	words := strings.Split(str, " ")
	fmt.Printf("%d", len(words))
}
