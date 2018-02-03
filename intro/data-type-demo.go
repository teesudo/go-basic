package main 

import "fmt"

func main() {
	//vars declare
	var str string
	var n, m int
	var mn float32

	//assign vals
	str = "Hello World"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("value of str = ", str)
	fmt.Println("value of n = ", n)
	fmt.Println("value of m = ", m)
	fmt.Println("value of mn = ", mn)

	//directly assign
	var city string = "New York"
	var x int = 100

	fmt.Println("value of city = ", city)
	fmt.Println("value of x = ", x)

	//declare vars with defining its type
	country := "DE"
	val := 15

	fmt.Println("value of country = ", country)
	fmt.Println("value of val = ", val)

	//define multiple variables
	var (
        name string
        email string
        age int
	)

	name = "john"
	email = "john@gmail.com"
	age = 27

	fmt.Println("value of name = ", name)
	fmt.Println("value of email = ", email)
	fmt.Println("value of age = ", age)

	//define var with vals
	var (
        location string = "Beijing"
        time int = 23
	)

	fmt.Println("value of location = ", location)
	fmt.Println("value of time = ", time)

}