package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)
	lower := strings.ToLower(str)
	if strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n") && strings.Contains(lower, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
