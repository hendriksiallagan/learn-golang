package main

import (
	"fmt"
	"strconv"
)

/*
	Create function from 1 to n parameter
	If n can be divided by 3 and 5, return "Fizzbuzz"
	If n can be divided by 3, return "Fizz"
	if n can be divided by 5, return "Buzz"
	else return n
*/

func FizzBuzz(n int) (res string) {

	res = ""
	for i:=1; i<=n; i++  {
		if i%15 == 0 {
			res += "FizzBuzz "
		}  else if i%3 == 0 {
			res+= "Fizz "
		} else if i%5 == 0 {
			res+="Buzz "
		} else {
			res += strconv.Itoa(i)
			res += " "
		} 	
	}

	return res
}

func main() {
	
	test_fb := FizzBuzz(15)

	fmt.Println(test_fb)
}
