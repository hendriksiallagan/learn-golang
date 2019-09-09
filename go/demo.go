package main

// you can also use imports, for example:
 import "fmt"
 //import "runtime"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// func Solution(A []int) (min int, max int) {
//     // write your code in Go 1.4
//     min = A[0]
//     max = A[0]
//     for _, value := range A {
//                  if value < min {
//                          min = value 
//                  }
                 
//                  if value > max {
//                      max = value
//                  }
//          }
         
//     return min, max
// }



func main() {
 
   //A := []int{1,3,6,4,1,2} //5
  //A := []int{-1,1,3}
  //A := []int{-1,-1,-3}

  A :=[]int{1,2,3,4,5}

   
   var temp [6]int 
   var B [6]int
   //var temp map[int]int{}
    end := len(A) - 1
    for {

        if end == 0 {
            break
        }

        for i := 0; i < len(A)-1; i++ {
            if A[i] > A[i+1] {
                A[i], A[i+1] = A[i+1], A[i]
            }

        }
        end -= 1
    }


  a :=0
  b :=0
  c:=0
  count:= len(A)-1

  result:=0

 for j:=0; j<count; j++ {
    if A[j] < 0 {
        b++
    }

    if b == count {
         result = 1
    } else {//-1 1 3
        if A[j]!=A[j+1] && A[j+1]!=A[j]+1 {
            temp[a] = A[j]+1
          
            a++

            for k:=0; k<len(temp); k++ {
                if temp[k] > 0 {
                    B[c]=temp[k]
                    c++
                }
            }
          
            min := B[0]

            for _, value := range B {
                     if value < min  {
                        if value > 0 {
                            min = value 
                        }
                             
                     }
                     
             }
             result = min
        } else {
            //fmt.Println(A[2])
            result = A[count]+1
        }

        
    }

 }
 fmt.Println(result)
 //fmt.Println(temp)
   
}
