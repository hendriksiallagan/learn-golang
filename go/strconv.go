package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("123", 4, 32)

	fmt.Println(a)
}