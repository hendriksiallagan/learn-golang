package main

import "fmt"

func Prime(n int) (res []int) {
	result := make([]int, 0)
	for bil:=1; bil<n; bil++ {
		i:=0
		for bag:=1; bag<n; bag++ {
			if bil%bag == 0 {
				i++ 
			}
		}
		if i==2 && bil!=1 {
			result = append(result, bil)
		}
	}

	return result
}


func main() {

	test := Prime(50)
	fmt.Println(test)
	
}