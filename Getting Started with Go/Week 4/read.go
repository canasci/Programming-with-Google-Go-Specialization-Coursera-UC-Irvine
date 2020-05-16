/*
PROMPT

Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file has a
first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for
the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read
each line of the text file and create a struct which contains the first and last names found in the
file. Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading all lines
from the file, your program should iterate through your slice of structs and print the first and last
names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	/* input with multiple strings
		scanner := bufio.NewScanner(os.Stdin)
	    scanner.Scan() // use `for scanner.Scan()` to keep reading
	    line := scanner.Text()
		fmt.Println("captured:",line)
	*/

	var txtfile string
	fmt.Printf("Please enter the name of the text file:")
	fmt.Scanln(&txtfile)

	type Person struct {
		fname string
		lname string
	}

	f, err := os.Open(txtfile)

	if err != nil {
		log.Fatal(err)
	}

	sli1 := make([]Person, 0, 1)
	arr := make([]string, 2)

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		arr = strings.Split(scanner.Text(), " ")
		sli1 = append(sli1, Person{
			arr[0],
			arr[1]})
	}
	f.Close()

	var i int

	for _, p := range sli1 {
		fmt.Printf("Person %d\nname:%v\nsurname: %v\n", i, p.fname, p.lname)
	}

}
