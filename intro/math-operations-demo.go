package main 

import "fmt"

func main() {
	//declare variable
	var a, b int
	a = 5
	b = 10

	/* arithmetic operation */
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
	fmt.Printf("%d / %d = %.2f\n", a, b, float32(a) / float32(b))
	fmt.Printf("%d * %d = %d\n", a, b, a * b)
}