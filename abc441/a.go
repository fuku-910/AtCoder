package main

import (
	"fmt"
)

func main() {
	var p, q, x, y int
	fmt.Scan(&p, &q, &x, &y)

	if p <= x && p + 100 >= x && q <= y && q + 100 >= y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}