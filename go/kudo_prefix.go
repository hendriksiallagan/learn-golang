package main

import (
	"fmt"
	"reflect"
	"strings"
)

func inArray2(val interface{}, array interface{}) (exists bool) {

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

func main() {
	result := make([]string, 0)

	strSplit := strings.Split("0811,0812,0813,0821,0822,0823,0851,0852,0853", ",")
	for _, v := range strSplit {
		result = append(result, v)
	}

	cek := inArray2("0812", result)

	fmt.Println(cek)

}
