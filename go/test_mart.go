package main

import (
	"fmt"
	"reflect"
)

func in_Array(val interface{}, array interface{})(exists bool) {

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


func Mart(A []string) []string {
	l := len(A)
	arr := make([]string, int(l))
	temp := make([]string, int(l))

	a:=0
	b:=0

	for _,v :=range A {
		if in_Array(v, temp) == true {
			arr[a]= v
			a++
		} else {
			temp[b]=v
			b++
		}
	}



	 return arr
}

func main() {
	croott := Mart([]string{"q", "w", "e", "w", "r", "e"})

	fmt.Println(croott)
}
