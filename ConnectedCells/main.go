
/*
	Find the length of connected cells of 1's in an matrix of '0' and '1'.
*/
package main

import "fmt"

var max = 0

func MaxOnes(b [][]byte,i int,j int) int {
	
	if i < 0 || j < 0 {
		return 0
	}
	if i >= len(b) || j >= len(b[0]) {
		return 0
	}
	
	if b[i][j] == 0 {
		return 0
	}
	
	b[i][j] = 0
	return 1 + MaxOnes(b,i+1,j) + MaxOnes(b,i-1,j)+ MaxOnes(b,i,j-1)+  MaxOnes(b,i,j+1)
	
}


func main() {
	v := [][]byte{{0,1,1,0,0,1,0,1,},{0,1,1,0,0,1,0,1,},{0,1,0,0,1,1,0,1,},{0,1,1,0,1,1,0,1,},}
	
	max := 0
	
	for i :=0;i<len(v);i++ {
		for j :=0;j<len(v[i]);j++ {
			v := MaxOnes(v,i,j)
			if v>max {
				max = v
			}
		}
	}
	
	fmt.Println( max )
}
