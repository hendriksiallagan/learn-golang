package main

import (
	"fmt"
	"reflect"
	"strings"
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

func Balloon(n string) int {
	balloon := map[string]int{}
	balloon["B"] = 0
	balloon["A"] = 0
	balloon["L"] = 0
	balloon["O"] = 0
	balloon["N"] = 0

	ball := []string{"B", "A", "L", "O", "N"}

	nUpper:= strings.ToUpper(n)
	s := strings.Split(nUpper, "")


	for _,v := range s {
		if inArray(v, ball) == true {
			balloon[v]++
		}
	}

	balloon["L"] = balloon["L"]/2
	balloon["O"] = balloon["O"]/2


	if balloon["B"]==balloon["A"] && balloon["A"]==balloon["L"] && balloon["L"]==balloon["O"] && balloon["O"]==balloon["N"] {
		return int(balloon["B"])
	} else {
		min_vl:=balloon["B"]
		for _, vl := range balloon {
			if min_vl > vl {
				min_vl = vl
			}
		}

		return int(min_vl)
	}
}

func main() {
	test:= Balloon("BALLOONBALL")

	fmt.Println(test)
}
