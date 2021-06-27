package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanf("%d\n%d", &x, &y)
	if y > 0 {
		if x > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	} else {
		if x > 0 {
			fmt.Println(4)
		} else {
			fmt.Println(3)
		}
	}
}
