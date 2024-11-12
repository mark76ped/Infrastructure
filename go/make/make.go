package make

import "fmt"

type floatMap map[string]float64

func (m floatMap)output() {
	fmt.Println(m)
}

func main () {
	userNames := make([]string, 2, 5) //2 is the inital spaces created, 5 is the max capacity

	userNames[0] = "Julie" //adding name to premade space 0
	userNames[1] = "John"
	userNames =append(userNames, "Max") //appending to the end of the map
	userNames =append(userNames, "Manuel")

	fmt.Println(userNames)


	//courseRatings := make(map[string]float64, 5) //preallocated memory for max 5 slots. no initial space needed or desired for map
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.5

	courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames{
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}