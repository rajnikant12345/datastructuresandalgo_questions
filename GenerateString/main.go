/*
Generate all string of length n, drawn from 0 to k-1
*/

package main

import (
	"fmt"
)

func KString(A []int, n int, k int) {
	if n < 1 {
		fmt.Println(A)
		return
	}
	for i := 0;i<k;i++ {
		A[n-1] = i 
		KString(A,n-1,k)
	}
}

func main() {
	l := 5
	A := make([]int,l)
	KString(A,l,l)
}
