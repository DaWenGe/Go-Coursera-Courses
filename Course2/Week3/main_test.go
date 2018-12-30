package main

import "testing"

func TestAnimal(t *testing.T) {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	testcases := []struct {
		animal    Animal
		operation string
		result    string
	}{
		{cow, "move", "walk"},
		{bird, "speak", "peep"},
		{snake, "eat", "mice"},
	}

	for _, testcase := range testcases {
		var result string
		switch testcase.operation {
		case "move":
			result = testcase.animal.Move()
		case "speak":
			result = testcase.animal.Speak()
		case "eat":
			result = testcase.animal.Eat()
		}

		t.Logf("Operation %s\nGot: %s\nExpected: %s\n", testcase.operation, result, testcase.result)
	}
}
