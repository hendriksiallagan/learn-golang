package main

//import "reflect"

//func inArray(val interface{}, array interface{}) (index [3]int) {
//	a:=0
//
//	switch reflect.TypeOf(array).Kind() {
//	case reflect.Slice:
//		s := reflect.ValueOf(array)
//
//		for i := 0; i < s.Len(); i++ {
//			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
//				index[a] = i
//				a++
//			}
//		}
//	}
//
//	return index
//}

func max() {
	var max_vl  = resL[0]

	for _, vl := range resL {
		if max_vl < vl {
			max_vl = vl
		}
	}
}
