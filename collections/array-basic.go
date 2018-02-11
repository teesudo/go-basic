package main 

import (
	"fmt"
	"math/rand"
)
func main() {

	//define array
	fmt.Println("Define Arrays")
	var numbers[5] int
	var cities[5] string
	var matrix[3][3] int

	//insert data
	fmt.Println(">>>>>insert array data")
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("Cities %d", i)
	}

	//insert matrics data
	fmt.Println(">>>>>insert matirx data")
	for i := 0; i < 3; i++ {
		for j:= 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	//display data
	fmt.Println(">>>>>display array data")
	fmt.Println(numbers)
	fmt.Println(cities)
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
		fmt.Println(cities[j])
	} 

	//display matrix data
	fmt.Println(">>>>>display matrix data")
	fmt.Println(matrix)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}