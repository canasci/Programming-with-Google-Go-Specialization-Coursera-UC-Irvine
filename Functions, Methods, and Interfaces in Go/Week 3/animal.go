/*
PROMPT

Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
	1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.

	The following table contains the three animals and their associated data which should be hard-coded into your program.

	Animal	Food eaten	Locomotion method	Spoken sound
	cow		grass		walk				moo
	bird	worms		fly					peep
	snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct declaration
type Animal struct {
	food,
	locomotion,
	noise string
}

// Eat method for type Animal
func (ani *Animal) Eat() {
	fmt.Println(ani.food)
}

// Move method for type Animal
func (ani *Animal) Move() {
	fmt.Println(ani.locomotion)
}

// Speak method for type Animal
func (ani *Animal) Speak() {
	fmt.Println(ani.noise)
}

// Init method for type Animal
func (ani *Animal) Init(val string) {
	if val == "cow" {
		ani.food = "grass"
		ani.locomotion = "walk"
		ani.noise = "moo"
	} else if val == "bird" {
		ani.food = "worms"
		ani.locomotion = "fly"
		ani.noise = "peep"
	} else if val == "snake" {
		ani.food = "mice"
		ani.locomotion = "slither"
		ani.noise = "hsss"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Please type an animal name and an action separated with a whitespace > ")
		input, _ := reader.ReadString('\n')

		var fields []string
		fields = strings.Split(input, " ")
		species := strings.TrimSpace(strings.ToLower(fields[0]))
		action := strings.TrimSpace(strings.ToLower(fields[1]))

		if species != "cow" && species != "bird" && species != "snake" {
			fmt.Printf("Please choose: cow, bird, or snake as animal.\n")
			continue
		}

		ani1 := Animal{}
		ani1.Init(species)

		switch action {
		case "eat":
			ani1.Eat()
		case "move":
			ani1.Move()
		case "speak":
			ani1.Speak()
		default:
			fmt.Printf("Please choose: eat, move, or speak as action. \n")
		}
	}
}
