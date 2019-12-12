package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ThreeSum(arr []int) (res bool) {
    first := arr[0]

    var result string

   result=""
  if len(arr) >= 4 {
     for i:=1; i<len(arr); i++ { 
       for j:=1; j<len(arr); j++ {
	   for k:=1; k<len(arr); k++{
		if arr[i] + arr[j] == first {
		    if strings.Contains(result, strconv.Itoa(arr[i])) == false && strings.Contains(result, strconv.Itoa(arr[j])) == false {
		    	result += "true "
	            }
		}
	    }	
       }
    }
}
    	

    if result == "" {
	res = false
    } else {
	res = true
    }

	return res
}

func main() {
	test := ThreeSum([]int {12, 3, 1, -5, -4, 7})
	
	fmt.Println(test)
	
}
