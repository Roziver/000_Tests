package main

import (
	"fmt"
)

func pointer(a *int) {
	*a *= 2
	fmt.Println(a)
	fmt.Println(&a)
}

func main() {
	number := 12
	fmt.Println(number)
	pointer(&number)
	fmt.Println(number)
}
