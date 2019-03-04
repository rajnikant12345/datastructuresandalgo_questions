package main

import (
	"fmt"
)

func Binary(A []int, n int) {
	if n == len(A) {
		fmt.Println(A);return
	}
	A[n] = 0 ; Binary(A,n+1) ; A[n] = 1 ; Binary(A,n+1)
}

func main() {
	A := make([]int,8)
	Binary(A,0)
}
