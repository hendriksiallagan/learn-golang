package main

import (
	"fmt"
	"reflect"
)

func inArray(val interface{}, array interface{})(exists bool) {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

func SeatingStudents(arr []int) (res int) {
	addNum1 := 0
	addNum2 := 0
	
  	first := arr[0]

	arr[0] = -1

  	res = 0

    for j:=1; j<=first; j++ {
      	if inArray(j, arr) == false {
			if j%2==1 {
				addNum1 = j+1
			} else if j%2==0 {
				addNum1 = j-1		
			}

			addNum2 = j+2
	
			if inArray(addNum1, arr) == false || inArray(addNum2, arr) == false {
				res++
			}
    	}	
	}

	return res

}

func main() {
	test := SeatingStudents([]int{8,1,8})
	
	fmt.Println(test)
}
