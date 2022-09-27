package main

import "fmt"

func main() {
	var s, a, b, c int
	fmt.Scan(&s)
	a = s / 100
	b = (s / 10) % 10
	c = s % 10
	if a != b && b != c && a != c {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
