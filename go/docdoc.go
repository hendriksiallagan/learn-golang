package main

import "fmt"
import "strconv"

func Fizzbuzz(n int)(res string) {
	res = ""

	for i:=1; i<=n; i++ {
		if i%15==0 {
			res += "FizzBuzz\n"
		} else if i%3==0 {
			res += "Fizz\n"
		} else if i%5==0 {
			res += "Buz\n"
		} else {
			res += strconv.Itoa(i)
			res += "\n"
		}
	}

	return res 
}

func main() {
		fizzbuzz := Fizzbuzz(100)

		fmt.Println(fizzbuzz)
}