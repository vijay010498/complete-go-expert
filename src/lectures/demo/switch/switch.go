package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap Item")
	case p < 10:
		fmt.Println("Moderately Priced item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy Siting")
	case Business:
		fmt.Println("Business Siting")
	case FirstClass:
		fmt.Println("FirstClass Siting")
	default:
		fmt.Println("Other siting")
	}
}
