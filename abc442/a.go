package main

import (
	"fmt"
)


func main() {
	var s string 
	var cnt int = 0
	fmt.Scanf("%s", &s)
	for _, i := range s {
		if i == 'j' || i == 'i' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
