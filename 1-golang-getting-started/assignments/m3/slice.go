package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("enter integers and stores the integers in a sorted slice")
	v := make([]int, 0, 3) // cap=3 would make sense rather than len=3
	var s string
	for {
		fmt.Print("enter an integer to be added to the slice: ")
		fmt.Scan(&s)
		if s == "X" {
			break
		}
        // convert string to int for input X letter
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
        // add into the slices
		v = append(v, n)
        // use qsort in official library
		sort.Ints(v)
        // print the result with default String method
		fmt.Println("the contents of the slice in sorted order:", v)
	}
}
