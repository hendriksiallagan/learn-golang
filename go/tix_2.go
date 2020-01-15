package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isFormat(n string) (res string) {
	s := strings.Split(n, "")
	temp := 0
	result := ""

	for i, v := range s {
		if v == "O" {
			temp++
		} else {
			temp = 0
		}

		temps := strconv.Itoa(temp)
		result += temps

		if i != len(s)-1 {
			result += "+"
		}
	}

	return result
}

func main() {
	n := isFormat("OOXXOXXOOO")

	fmt.Println(n)
}
