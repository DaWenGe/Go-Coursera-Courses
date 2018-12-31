package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func partialSort(wg *sync.WaitGroup, slice []int) {
	fmt.Println("slice to be sorted: ", slice)
	sort.Ints(slice)
	fmt.Println("sorted slice: ", slice)
	wg.Done()
}

func mergeSort(left []int, right []int) {
	i := 0
	j := 0
	lenLeft := len(left)
	lenRight := len(right)
	var result []int
	for i < lenLeft && j < lenRight {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}

	if i < lenLeft {
		result = append(result, left[i:]...)
	}

	if j < lenRight {
		result = append(result, right[j:]...)
	}

	copy(left, result[:lenLeft])
	copy(right, result[lenLeft:])
}

func main() {
	fmt.Println("Please enter a sequence of integers separated by space:")
	array := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		for _, v := range strs {
			number, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error parsing integer.")
				continue
			}
			array = append(array, number)
		}
	}

	fmt.Println(array)
	length := len(array)
	if length < 4 {
		fmt.Fprintf(os.Stderr, "please make sure you enter at least 4 integers\n")
		return
	}
	segmentSize := length / 4

	subArrayOne := array[0:segmentSize]
	subArrayTwo := array[segmentSize : segmentSize*2]
	subArrayThree := array[segmentSize*2 : segmentSize*3]
	subArrayFour := array[segmentSize*3:]

	var wg sync.WaitGroup
	wg.Add(4)

	go partialSort(&wg, subArrayOne)
	go partialSort(&wg, subArrayTwo)
	go partialSort(&wg, subArrayThree)
	go partialSort(&wg, subArrayFour)

	wg.Wait()

	mergeSort(subArrayOne, subArrayTwo)
	mergeSort(subArrayThree, subArrayFour)
	mergeSort(array[:segmentSize*2], array[segmentSize*2:])

	fmt.Println("Sorted array: ", array)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
