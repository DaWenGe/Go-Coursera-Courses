package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, pos int) {
	tmp := slice[pos]
	slice[pos] = slice[pos-1]
	slice[pos-1] = tmp
}

func BubbleSort(slice []int) {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := 1; j < length-i; j++ {
			if slice[j] < slice[j-1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	fmt.Println("Please enter 10 integers separated by spaces:")

	numbers := make([]int, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		for _, v := range strs {
			number, err := strconv.Atoi(v)
			if err == nil {
				numbers = append(numbers, number)
			}
		}
	}

	fmt.Println("numbers: ", numbers)
	BubbleSort(numbers)
	fmt.Println("sorted: ", numbers)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
