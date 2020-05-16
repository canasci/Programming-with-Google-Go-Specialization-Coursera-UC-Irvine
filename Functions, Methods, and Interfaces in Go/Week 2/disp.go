/*
PROMPT

Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s = 1/2 * a * t^2 + v0 *t + s0

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given
 values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t,
 representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	Fn := func(t float64) float64 {
		return (0.5*a*math.Pow(t, 2) + v0*t + s0)
	}
	return Fn
}
func main() {
	var str1 string
	var arr []float64

	fmt.Printf("\tPlease enter the values for acceleration, initial velocity,\nand initial displacement separately:\n")

	for i := 0; i < 3; i++ {
		fmt.Scanln(&str1)
		f, _ := strconv.ParseFloat(str1, 64)
		arr = append(arr, f)
	}

	var time float64
	fmt.Printf("Please enter the time passed: ")
	fmt.Scanln(&time)

	fn := genDisplaceFn(arr[0], arr[1], arr[2])

	fmt.Printf("Displacement after %.3f seconds is %.3f m\n", time, fn(time))
}
