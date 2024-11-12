package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	fmt.Println(websites) //prints all contents of map
	fmt.Println(websites["Google"])  //prints from key map value
	websites["LinkedIn"] = "https://linkedin.com"   //adds item to map
	fmt.Println(websites["LinkedIn"])

	delete(websites, "Google") //deletes from map via key
	fmt.Println(websites)
}