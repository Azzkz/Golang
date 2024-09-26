package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	fmt.Println("Sum:", add(3, 4))
	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b)
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
