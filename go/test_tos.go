package main

import "fmt"

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func Tree(A []int)int {
	plus:=0
	minus:=0
	for i:=0; i<len(A)-1; i++ {
		if A[i]<A[i+1]{
			plus++
		} else if A[i]>A[i+1] {
				minus++
		}
	}

	//if A[len(A)-2] < A[len(A)-1] {
	//	plus++
	//} else  if A[len(A)-2] > A[len(A)-1] {
	//	minus++
	//}

	//N:= RemoveIndex(A, 1)
	//fmt.Println(N)

	//N:=make([]int, len(A)-1)
	//var N map[int]int


	if plus==len(A)-1 || minus==len(A)-1 {
		return -1
	} else if plus==minus+1 || plus==minus-1 {
		return 0
	} else {
		//for j:=0; j<len(A)-1; j++ {
		//	N[j]= RemoveIndex(A, j)
		//
		//	fmt.Println(N[j])
		//}
		if plus > minus {
			return plus
		} else {
			return minus
		}

	}


	//fmt.Println(len(A))

	//fmt.Println(plus)
	//fmt.Println(minus)


}

func main() {
	//as:=Tree([]int{1,2,3,4})
	//as:=Tree([]int{1,3,1,2})
	as:=Tree([]int{3,4,5,3,7})

	fmt.Println(as)
}

