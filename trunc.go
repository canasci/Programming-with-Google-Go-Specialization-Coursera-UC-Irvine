/*
PROMPT

Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.

Truncation is the process of removing the digits to the right of the decimal place.
*/
package main

import "fmt"

func main() {

	var userfloat float32

	fmt.Printf("Hello user, could you please enter a floating point?")

	fmt.Scan(&userfloat)

	fmt.Printf("Truncated version of %f is %.0f", userfloat, userfloat)

}
