package main

import (
	"fmt"
)

func main() {
	var p, q, x, y int
	fmt.Scan(&p, &q, &x, &y)

	p_100 := p + 99
	q_100 := q + 99

	if p <= x && p_100 >= x && q <= y && q_100 >= y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}