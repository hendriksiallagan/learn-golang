package main

import "fmt"

func Solution(A []int) int {
	var result int = 0;
	var i int;
	fmt.Println(len(A))
	for i = 0; i < len(A); i++ {
		if result > A[i]  {
			result = A[i];
		}
	}
	return result;
}

func main() {
	a := Solution([]int{-1000,1000,-2000,2000})

	fmt.Println(a)
}
