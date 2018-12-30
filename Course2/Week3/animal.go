package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {
	animals := map[string]*Animal{
		"cow":   &Animal{"grass", "walk", "moo"},
		"bird":  &Animal{"worms", "fly", "peep"},
		"snake": &Animal{"mice", "slither", "hsss"},
	}

	fmt.Println("Please enter the animal name and operation: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		var result string

		switch strs[1] {
		case "eat":
			result = animals[strs[0]].Eat()
		case "move":
			result = animals[strs[0]].Move()
		case "speak":
			result = animals[strs[0]].Speak()
		}

		fmt.Printf("Requested animal: %s - Requsted operation: %s\n", strs[0], strs[1])
		fmt.Println("Response: ", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
