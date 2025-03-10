
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(9)
	fmt.Println("The magic number is", rand.Intn(10))
}