package main

import (
	."fmt"
)

func main() {
	var n int = 0
	Scan(&n)
	lista := make([]int, n)
	Printf("Inserire %d numeri", n)
	for i := 0; i < n; i++ {
		var numero int
		Scan(&numero)
		lista[(n-1)-i] = numero 
	}
	Println(lista)
}