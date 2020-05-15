/*
PROMPT
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/

package main

import "fmt"

func div(v1 int, v2 int) int {
	return v1 / v2
}

func alterVar(p1 *int) {
	*p1 = (*p1) * 2
}

func main() {
	a := 20
	b := 10
	go alterVar(&a)        // If executes, it will alter variable "a" to be two times its previous value
	fmt.Println(div(a, b)) // If there is no race, console will print 4; else it prints 2
}

/*
Race conditions may occur, due to communication, when outcome depends on non-deterministic ordering of instructions of separate goroutines(main is also a goroutine).
All the possible machine code level interleavings must be considered. To prevent data races, one can synchronize concurrent routines by utilizing waitgroups(for main goroutine to wait
for subgoroutines) and channels.

In this program, we have a main routine and a function call as a separate goroutine inside main. Data race occurs here since the main goroutine finishes before
the other goroutine executes successfully. We can use the Go built-in race detector to detect potential data race conditions.

To use Go race detector: Open terminal(powershell, command prompt etc.) and go to the directory which contains the main program to be investigated.
Then type "go run -race" followed by a whitespace and the main file name with extension(i.e "race.go" in this case) to run the program with race condition detection.
The "-race" flag can also be added to the "go test" and "go build" commands.

The following command prompt lines are the race condition analysis for this program.
The first line which prints "2" is the result of the program.

Below, we can see that the information about our program's data race is divided into three sections:
[1] The first section tells us there was an attempted write to inside a goroutine that we created
[2] This section tells us there was a simultaneous read by the main goroutine, which in our code, traces through the print statement included in the main goroutine
[3] The third section conveys the information where the goroutine that caused section [1] was created
P.S. If there is no race condition, race condition detector prints nothing: If "-race" flag is added to "go run", terminal only gives the output of the program executed.

PS C:\Users\canas\OneDrive\Belgeler\Golang Coursera\Concurrency in Go\Week 2> go run -race race.go
2
==================
WARNING: DATA RACE
Write at 0x00c000122068 by goroutine 7:
  main.alterVar()
      C:/Users/canas/OneDrive/Belgeler/Golang Coursera/Concurrency in Go/Week 2/race.go:10 +0x57

Previous read at 0x00c000122068 by main goroutine:
  main.main()
      C:/Users/canas/OneDrive/Belgeler/Golang Coursera/Concurrency in Go/Week 2/race.go:17 +0x8f

Goroutine 7 (running) created at:
  main.main()
      C:/Users/canas/OneDrive/Belgeler/Golang Coursera/Concurrency in Go/Week 2/race.go:16 +0x81
==================
Found 1 data race(s)
exit status 66
*/
