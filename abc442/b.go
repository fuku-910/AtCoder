package main

import (
	"fmt"
)

func main() {
	var q int
	var vol int = 0
	var play bool = false
	
	fmt.Scan(&q)

	for i := 0; i <q; i++ {
		var a int
		fmt.Scan(&a)
		
		switch a {
		case 1:
			vol++
		case 2:
			if vol > 0 {
				vol--
			}
		case 3:
			play = !play
		}
		if play && vol > 2{
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
} 
