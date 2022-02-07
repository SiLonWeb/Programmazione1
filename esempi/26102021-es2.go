package main

import . "fmt"

func main()  {
	var n int
	Scan(&n)

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			for r := 1; r <= n*n; r++ {
				if r != a*b && r%9 == (a*b)%9{
					Println("")
				}
			}			
		}
	}
}