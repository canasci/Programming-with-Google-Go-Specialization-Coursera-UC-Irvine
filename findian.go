/*
PROMPT

Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.package main
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var str1 string

	fmt.Printf("Hello user, could you please enter a name?")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str1 = scanner.Text()

		if strings.HasPrefix(strings.ToLower(str1), "i") && strings.Contains(strings.ToLower(str1), "a") && strings.ToLower(str1[len(str1)-1:]) == "n" {
			fmt.Printf("Found!")
		} else {
			fmt.Printf("Not Found!")
		}
	}

	var idMap map[string]int
	idMap["a"] = 1
	for key, val := range idMap {
		fmt.Println(key, val)
	}
	type Person struct {
		name string
	}
	var p1 Person
	p1.name = "Can"

}
