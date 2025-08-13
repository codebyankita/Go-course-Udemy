package intermediate

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)
	// n * (n - 1) * (n-2) * factorial (n-3)..... factorial(0)

}

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}