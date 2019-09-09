package main

import (
	"fmt"
	"reflect"
)

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func Solution(A []int, K int, L int) int {

	count := int(K)+int(L)
	lenArr := len(A)
	var result int

	//var resultL int
	//var max int
	//var temp [5]int

	templ := make([]int, 0)

	//j:=0

	if count > lenArr {
		result = -1
	} else {
		result = 0

		la:=0
		z:=0
		for i:=0; i<lenArr; i++ {
			if la <= L {
				for l:=0; l<L; l++ {
					templ[z] = A[i] + A[i+1]
				}
			} else {
				la = 0
			}
		}

		maxl := templ[0]

		for _, value := range templ {
			if value > maxl {
				maxl = value
			}
		}





		//end := len(A) - 1
		//for {
		//	if end == 0 {
		//		break
		//	}
		//	for i := 0; i < len(A)-1; i++ {
		//		if A[i] < A[i+1] {
		//			A[i], A[i+1] = A[i+1], A[i]
		//		}
		//
		//	}
		//	end -= 1
		//}
		//
		//fmt.Println(lenArr)
		//fmt.Println(A)
		//fmt.Println(A[7])

		//for i:=0; i<lenArr; i++ {
		//	bool, _ := in_array(A[i], temp)
		//	if bool == false && j<=count {
		//		temp[j] = A[i]
		//		j++
		//	}
		//	temp[0] = 5
		//}
		//
		//temp[0] = 5
		//
		//
		//
		//for k:=0; k<len(temp); k++ {
		//	result = result + temp[k]
		//}



	}


	//fmt.Println(count)
	//fmt.Println(lenArr)
	return result
}


func main() {
	a := Solution([]int{6,1,4,6,3,2,7,4}, 3, 2)
	//a := Solution([]int{6,1,4}, 3, 2)

	//A:=[]int{6,1,4}
	//
	//end := len(A) - 1
	//for {
	//	if end == 0 {
	//		break
	//	}
	//	for i := 0; i < len(A)-1; i++ {
	//		if A[i] < A[i+1] {
	//			A[i], A[i+1] = A[i+1], A[i]
	//		}
	//
	//	}
	//	end -= 1
	//}

	fmt.Println(a)
	//fmt.Println(A)
}