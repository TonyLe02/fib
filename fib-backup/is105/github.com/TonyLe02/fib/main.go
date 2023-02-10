package main

import (
	"fmt"

	"github.com/TonyLe02/fib"
)

func main() {
	var input int
	fmt.Println("Sett inn tall:")
	fmt.Scan(&input)
	fmt.Println(fib.Fib(input))
}
