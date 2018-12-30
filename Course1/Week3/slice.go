package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		var number int
		fmt.Println("Please enter an integer: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Entered some non integer, program will quit...")
			break
		}
		slice = append(slice, number)
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
		fmt.Println(slice)
	}
}
