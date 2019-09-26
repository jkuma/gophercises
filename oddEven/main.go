package main

import "fmt"

func main() {
	integerSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, integer := range integerSlice {
		if (integer % 2) == 0 {
			fmt.Println(integer, "is even")
			continue
		}
		fmt.Println(integer, "is odd")
	}
}
