package main

import "fmt"

func main() {
	x := []int{5,4,8,10,3,7,2,1}
	
	end := len(x)-1
	
	for {
		if end == 0 {
			break
		}
		
		for i:=0; i<len(x)-1; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
			}
		}
		
		end -=1
		
	}
	
	fmt.Println(x)
}