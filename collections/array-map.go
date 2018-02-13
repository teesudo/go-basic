package main 

import (
	"fmt"
	"math/rand"
)

func main() {

	// define array
	fmt.Println("define map")
	products := make(map [string] int)
	customers := make(map [string] int)

	//insert data
	fmt.Println(">>>>>insert map data")
	products["product1"] = rand.Intn(100)
	products["product2"] = rand.Intn(100)

	customers["cust1"] = rand.Intn(100)
	customers["cust2"] = rand.Intn(100)

	//display data
	fmt.Println(products)
	fmt.Println(customers)

}