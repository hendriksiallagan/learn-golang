package main

import (
	"fmt"
)

/**
Create a function to search the second biggest number of array
if length of array less than 2 or more than 1024 return -1
if number sin array are same return -1
*/


func FormulaSecondBiggest(A []int) (s int) {
	var first int
	var second int 


	if len(A) < 2 || len(A)>1024 {
		return -1
	}

	int_min := A[0]

	first = int_min
	second = int_min

	for _,v := range A {
		
		if v>=first {
			second = first
			first = v
		} else {		
			if second==first || v >= second {
				second = v
			}
		}
	}

	if first==second {
		return -1
	}

	return second
}

func main() {
	result := FormulaSecondBiggest([]int{5, 5, 4, 2})
	//result := FormulaSecondBiggest([]int{3, -2})
	//result := FormulaSecondBiggest([]int{4, 4, 4, 4})
	//result := FormulaSecondBiggest([]int{})
	
	
	fmt.Println(result)
	
}