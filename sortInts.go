/*
PROMPT:

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.

p.s. This program utilizes bubblesort algorithm in order to sort the integer series
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func bubbleSort(sli []int) []int {
	if len(sli) == 1 {
		return sli
	}

	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
	return sli
}

func swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func readSeriesInt() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please type a series of integers as comma separated values> ")
	input, _ := reader.ReadString('\n')
	var seriesStr []string
	seriesStr = strings.Split(input, ",")
	var seriesInt []int

	for i := range seriesStr {
		seriesStr[i] = strings.TrimSpace(seriesStr[i])
		num, _ := strconv.Atoi(seriesStr[i])
		seriesInt = append(seriesInt, num)
	}
	return seriesInt
}

func sliceSort(sli []int, wg *sync.WaitGroup, c chan []int) {
	fmt.Printf("%d is to be sorted\n", sli)
	sli = bubbleSort(sli)
	c <- sli
	wg.Done()
}

func main() {
	var unsortedSliceInt []int
	unsortedSliceInt = readSeriesInt()
	fmt.Printf("The unsorted slice: %d\n", unsortedSliceInt)

	var wg sync.WaitGroup

	size1 := len(unsortedSliceInt) / 4

	c1 := make(chan []int, 4)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		if i < 3 {
			go sliceSort(unsortedSliceInt[i*size1:(i+1)*size1], &wg, c1)
		} else if i == 3 {
			go sliceSort(unsortedSliceInt[i*size1:], &wg, c1)
		}
	}

	wg.Wait()
	close(c1)

	var mergedArray []int // initialized as slice, not array since there is no limitation to the input
	for slice := range c1 {
		for _, val := range slice {
			mergedArray = append(mergedArray, val)
		}
	}

	fmt.Printf("Merged array formed by the slices sorted by separate goroutines: %d\n", mergedArray)

	mergedArray = bubbleSort(mergedArray)
	fmt.Printf("Sorted merged array: %d\n", mergedArray)
}
