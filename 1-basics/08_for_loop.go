package main

import "fmt"

func main() {

	// Simple iteration over a range
	// for initialization;condition;post{code block to be executed repeatedly}
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %v, Value:%v\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // continue the loop but skip the 
			//rest of lines/statements
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break // break out of the loop
		}
	}

	// ASTERISK LAYOUT
	rows := 5

	//Outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line
	}

	// Go 1.22 update
	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")

}