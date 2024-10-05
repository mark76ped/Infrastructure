// package main

// import "fmt"

// func main() {
// 	prices := []float64{
// 		10.99,
// 		8.99,
// 	}

// 	fmt.Println(prices[1])
// 	fmt.Println(prices[0])

// 	prices = append(prices, 5.99) //adds element to array
// 	fmt.Println(prices)
// 	prices = prices[1:] //removes element from array
// 	fmt.Println(prices)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{
// 		10.99,
// 		9.99,
// 		45.99,
// 		20.0,
// 	}
// 	productNames[2] = "Ultralearning"
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:] //slices array
// 	featuredPrices[0] = 199.99 //updating a slice, which updates the original array
// 	highlightedPrices := featuredPrices[:1] //gets the index 1 of a slice (slicing a slice)
// 	fmt.Println(featuredPrices) //slice
// 	fmt.Println(highlightedPrices) //slice of slice
// 	fmt.Println(prices) // shows updated array even though modification was done on slice
// 	fmt.Println(prices[(len(prices)-1):]) //gets last element in array
// 	fmt.Println(cap(highlightedPrices)) //outputs 3 since highlightedPrices is a slice of featuredPrices which has 3 value
// }