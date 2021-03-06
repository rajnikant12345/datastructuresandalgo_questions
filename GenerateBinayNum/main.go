/*
Genarate all binary strings of bit n
*/

package main

import (
	"fmt"
)
// Recursive funtion for generating binary numbers
func Binary(A []int, n int) {
	// if index is out of bound print and return
	if n == len(A) {
		fmt.Println(A);return
	}
	// set current index to 0 and recurse
	A[n] = 0 ; Binary(A,n+1) 
	// set current index to 1 and recurse
	A[n] = 1 ; Binary(A,n+1)
}

func main() {
	A := make([]int,8)
	Binary(A,0)
}
