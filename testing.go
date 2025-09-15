package main 

import (
	"fmt"
)

func main() {
	a:= 1
	b:= 3
	fmt.Println("the sum is", Sum(a,b))
}

func Sum(a,b int) int {
	return a+b
}