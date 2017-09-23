package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(2524543243565456390)
	fmt.Println("My favorite number is", rand.Intn(10))
}
