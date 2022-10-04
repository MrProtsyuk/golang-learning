package main

import "fmt"

// algorithm to sort through numbers 0-10 to see what is even or odd
func main() {
	alphanumeric := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range alphanumeric {
		if  value%2 == 0 {
			fmt.Println(value," is even")
		} else {
			fmt.Println(value," is odd")
		}
	}

}