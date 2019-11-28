package main

import (
	"fmt"
)

func Palindrome(n string)(res bool) {
	for i:=0; i<len(n); i++ {
		if n[i] != n[len(n)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	test := Palindrome("anna")
	fmt.Println(test)
}