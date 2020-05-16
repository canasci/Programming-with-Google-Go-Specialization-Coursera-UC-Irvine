/*
PROMPT
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user.

Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.

The following table contains the three types of animals and their associated data.

	Animal	Food eaten	Locomtion method	Spoken sound
	cow		grass		walk				moo
	bird	worms		fly					peep
	snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface declaration
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type declaration
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// Eat method for type Cow
func (c Cow) Eat() {
	fmt.Println(c.food)
}

// Move method for type Cow
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

// Speak method for type Cow
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// Bird type declaration
type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// Eat method for type Bird
func (b Bird) Eat() {
	fmt.Println(b.food)
}

// Move method for type Bird
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

// Speak method for type Bird
func (b Bird) Speak() {
	fmt.Println(b.noise)
}

// Snake type declaration
type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// Eat method for type Snake
func (s Snake) Eat() {
	fmt.Println(s.food)
}

// Move method for type Bird
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

// Speak method for type Bird
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	animalMap := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Please type a command to create or query an animal > ")
		input, err := reader.ReadString('\n')

		if err == nil {
			var words []string
			words = strings.Split(input, " ")
			command := strings.TrimSpace(strings.ToLower(words[0]))

			if command == "newanimal" {
				animalName := strings.TrimSpace(strings.ToLower(words[1]))
				animalType := strings.TrimSpace(strings.ToLower(words[2]))
				var ani Animal
				if animalType == "cow" {
					ani = Cow{animalName, "grass", "walk", "moo"}
					animalMap[animalName] = ani
				} else if animalType == "bird" {
					ani = Bird{animalName, "worms", "fly", "peep"}
					animalMap[animalName] = ani
				} else if animalType == "snake" {
					ani = Snake{animalName, "mice", "slither", "hsss"}
					animalMap[animalName] = ani
				}
				fmt.Printf("Created it!\n %s: %+v\n", animalType, ani)
			} else if command == "query" {
				animalName := strings.TrimSpace(strings.ToLower(words[1]))
				action := strings.TrimSpace(strings.ToLower(words[2]))
				if action == "eat" {
					animalMap[animalName].Eat()
				} else if action == "move" {
					animalMap[animalName].Move()
				} else if action == "speak" {
					animalMap[animalName].Speak()
				}
			} else {
				fmt.Printf("Invalid command! Please try a valid one\n")
			}
		}
	}
}
