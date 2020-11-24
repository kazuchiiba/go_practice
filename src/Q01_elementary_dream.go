package main

import (
	"fmt"
)

/*
	https://yukicoder.me/problems/no/1249
	2020/11/24
*/

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b == c {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
