package main

import "fmt"

func BubbleSort(s []int) {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func main() {
	fmt.Print("Enter a sequence of up to 10 integers: ")
	s := make([]int, 10)
	for i := range s {
		fmt.Scan(&s[i])
	}
	BubbleSort(s)
	fmt.Println("Sorted:", s)
}
