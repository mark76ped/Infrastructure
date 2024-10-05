package main

import (
	"fmt"
)

type Product struct{
	id string
	title string
	price float64
}
func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{
		"Hiking",
		"Reading",
		"Learning",
	}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	selectHobbies := hobbies[1:]
	fmt.Println(selectHobbies)
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	firstHobbies := hobbies[:2]
	fmt.Println(firstHobbies)
	// firstHobbiestoo := hobbies[0:2]
	// fmt.Println(firstHobbiestoo)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	firstAndLast := firstHobbies[1:3]
	fmt.Println(firstAndLast)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{
		"learn Golang well",
		"stay employed",
	}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "get promoted"
	courseGoals = append(courseGoals, "finally enjoy life")
	fmt.Println(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{"first-product", "Shampoo", 9.99},
		{"second-product", "Apples", 1.99},
	}
	fmt.Println(products)
	newProduct := Product {"third-product", "A Golang Book", 54.99}
	products = append(products, newProduct)
	fmt.Println(products)
}
