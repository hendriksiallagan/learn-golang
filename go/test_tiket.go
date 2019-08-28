package main

import (
	"fmt"
)


func Sol(A []int, K int, L int) int {
	var result int
	if K+L > len(A) {
		result = -1
	} else {
		k:=0
		l:=0
		resK := make([]int, int(len(A)-K)+1)
		resL := make([]int, int(len(A)-L)+1)

		//count array per K
		for i:=0; i<=int(len(A)-K); i++ {
			resK[k] = 0

			for j:=i; j<K+i; j++ {
				resK[k] = resK[k] + A[j]
			}
			k++
		}

		//find max value of resL
		var max_vk  = resK[0]

		for _, vk := range resK {
			if max_vk < vk {
				max_vk = vk
			}
		}

		//count array per L
		for n:=0; n<=int(len(A)-L); n++ {
			resL[l] = 0

			for m:=l; m<L+n; m++ {
				resL[l] = resL[l] + A[m]
			}
			l++
		}

		//find max value of resL
		var max_vl  = resL[0]

		for _, vl := range resL {
			if max_vl < vl {
				max_vl = vl
			}
		}


		result = int(max_vk) + int(max_vl)

	}

	return result

}


func main() {

	a := Sol([]int{6,1,4,6,3,2,7,4}, 3, 2)

	fmt.Println(a)

}