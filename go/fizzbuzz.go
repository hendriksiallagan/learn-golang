package main

import (
	"fmt"
)

func main() {
	
	for i:=0; i<20; i++  {
		if i%15 == 0 {
		  fmt.Println("Fizz Buzz")
		}  else if i%3 == 0 {
	  	 fmt.Println("Fizz")
		} else if i%5 == 0 {
		  fmt.Println("Buzz")
		} else {
		 fmt.Println(i)
		}
	}
}