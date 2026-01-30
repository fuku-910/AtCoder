package main

import (
	"fmt"
)

func main() {
	var q int
	var aArr []int
	fmt.Scan(&q)

	for i := 0; i <q; i++ {
		var a int
		fmt.Scan(&a)
		aArr = append(aArr, a)
	}


	var vol int = 0
	var play bool = false
	for _, num := range aArr {
		switch num {
		case 1:
			vol++
		case 2:
			if vol > 0 {
				vol--
			}
		case 3:
			play = !play
		}
		if play == true && vol > 2{
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} 
}