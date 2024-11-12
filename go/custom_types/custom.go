package custom

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "Max"

	name.log()
}