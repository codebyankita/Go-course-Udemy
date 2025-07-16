package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	// underscore is blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Printf(" Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray1 := [3]int{1, 2, 3}
	var copiedArray1 *[3]int

	copiedArray1 = &originalArray1
	copiedArray1[0] = 100

	fmt.Println("Original array1:", originalArray1)
	// fmt.Println("Copied array:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
