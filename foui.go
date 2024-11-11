package main

import "fmt"

func main() {
	fmt.Println("in:")
	fmt.Scan()
	var a string
	var b string
	b = "wo"
	fmt.Scan(&a)
	if a == b {
		fmt.Println("欢迎")
		return
	} else {
		fmt.Println(a)
		for l := 1; l <= 999999999; l++ {
			fmt.Println(l)
		}
	}

	return
}
