package main

// "factored" import statement.
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Hello World, my luck number is %d", rand.Intn(10))
}
