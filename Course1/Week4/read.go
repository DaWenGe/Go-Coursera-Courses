package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	FName string
	LName string
}

func main() {
	var names []Name
	var fileName string

	fmt.Println("Please enter your file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		names = append(names, Name{strs[0], strs[1]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for i, v := range names {
		fmt.Println("Name at index: ", i)
		fmt.Println("FName: ", v.FName)
		fmt.Println("LName: ", v.LName)
	}

	file.Close()
}
