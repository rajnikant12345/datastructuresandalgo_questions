/*
Given a string with 0 and 1 and ?. Generate all the possible string just filling ? 
*/

package main

import "fmt"

func GetQueryArray(in string) []int {
	a := make([]int,0)
	for idx,i := range in {
		if i == '?' {
			a = append(a,idx)	
		}
	}
	return a
}

func Kstring(in []byte,a []int,n int) {
	if n<0 {
		fmt.Println(string(in))
		return
	}
	idx := a[n]
	in[idx] = byte('0')
	Kstring(in,a,n-1)
	idx = a[n]
	in[idx] = byte('1')
	Kstring(in,a,n-1)
	
}


func main() {
	in := "01?0?101"
	f := GetQueryArray(in)
	Kstring([]byte(in),f,len(f)-1)
}
