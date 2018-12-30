package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(acc, velocity, displacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*acc*time*time + velocity*time + displacement
	}

	return fn
}

func main() {
	fmt.Println("Please enter values for acceleration, initial velocity and initial displacement: ")

	var fn func(float64) float64
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()

		var acc, velocity, displacement float64
		for i, v := range strings.Fields(line) {
			if value, err := strconv.ParseFloat(v, 64); err == nil {
				switch i {
				case 0:
					acc = value
				case 1:
					velocity = value
				case 2:
					displacement = value
				}
			}
		}

		fn = GenDisplaceFn(acc, velocity, displacement)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println("Please enter the time: ")
	var time float64
	fmt.Scan(&time)

	fmt.Println("Displacement travelled: ", fn(time))
}
