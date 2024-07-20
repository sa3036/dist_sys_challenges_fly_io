package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// := short assignment statement can be used in place of a var declaration with implicit type.
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
