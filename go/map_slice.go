package main

import "fmt"

func main() {

	//map: example 1
	var month1 map[string]int

	month1 = map[string]int{}

	month1["january"] = 1
	month1["february"] = 2

	//map: example 2
	var month2 = map[string]int{
		"maret":3,
		"april":4,
	}

	fmt.Println(month1)
	fmt.Println(month2)
}