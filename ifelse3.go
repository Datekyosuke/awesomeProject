package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 10 {
		fmt.Print(a)
	} else if a >= 10 && a < 100 {
		fmt.Print(a / 10)
	} else if a >= 100 && a < 1000 {
		fmt.Print(a / 100)
	} else if a >= 1000 && a < 10000 {
		fmt.Print(a / 1000)
	} else if a == 10000 {
		fmt.Print("1")
	}
}
