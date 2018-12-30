package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (cow *Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow *Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow *Cow) Speak() {
	fmt.Println(cow.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (bird *Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird *Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird *Bird) Speak() {
	fmt.Println(bird.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (snake *Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake *Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake *Snake) Speak() {
	fmt.Println(snake.noise)
}

func animalFactory(animalType string) Animal {
	var animal Animal
	switch animalType {
	case "cow":
		animal = &Cow{"grass", "walk", "moo"}
	case "bird":
		animal = &Bird{"worms", "fly", "peep"}
	case "snake":
		animal = &Snake{"mice", "slither", "hsss"}
	}

	return animal
}

func queryAnimal(operation string, animal Animal) {
	if animal == nil {
		fmt.Println("invalid animal")
		return
	}
	switch operation {
	case "eat":
		animal.Eat()
	case "speak":
		animal.Speak()
	case "move":
		animal.Move()
	}
}

func main() {
	animals := make(map[string]Animal)
	fmt.Println(">")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)

		switch strs[0] {
		case "newanimal":
			animals[strs[1]] = animalFactory(strs[2])
			fmt.Println("Created it!")
		case "query":
			animal := animals[strs[1]]
			queryAnimal(strs[2], animal)
		default:
			fmt.Println("Error parsing command")
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
